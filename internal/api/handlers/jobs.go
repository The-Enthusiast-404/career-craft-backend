// jobs.go

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"dev.theenthusiast.career-craft/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
)

type JobHandler struct {
	DB *sqlx.DB
}

func NewJobHandler(db *sqlx.DB) *JobHandler {
	return &JobHandler{DB: db}
}

func (h *JobHandler) GetJobsByCompany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	company := ps.ByName("company")

	var jobs []models.Job
	err := h.DB.Select(&jobs, "SELECT * FROM jobs WHERE company = $1", company)
	if err != nil {
		http.Error(w, "Failed to fetch jobs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}

func (h *JobHandler) CreateJob(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var job models.Job
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO jobs (title, company, url)
              VALUES ($1, $2, $3) RETURNING id`

	err = h.DB.QueryRowx(query, job.Title, job.Company, job.URL).Scan(&job.ID)
	if err != nil {
		http.Error(w, "Failed to create job", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(job)
}

func (h *JobHandler) BulkCreateJobs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var jobs []models.Job
	err := json.NewDecoder(r.Body).Decode(&jobs)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	fmt.Printf("Received %d jobs for bulk creation\n", len(jobs))

	tx, err := h.DB.Beginx()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to start transaction: %v", err), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	stmt, err := tx.Preparex(`
		INSERT INTO jobs (title, company, url)
		VALUES ($1, $2, $3)
		ON CONFLICT (company, title) DO UPDATE SET
		url = EXCLUDED.url
	`)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to prepare statement: %v", err), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	inserted := 0
	updated := 0
	for i, job := range jobs {
		if job.Title == "" || job.Company == "" {
			fmt.Printf("Skipping job %d due to missing required fields: %+v\n", i, job)
			continue
		}

		result, err := stmt.Exec(job.Title, job.Company, job.URL)
		if err != nil {
			fmt.Printf("Error inserting/updating job %d: %v\nJob details: %+v\n", i, err, job)
			continue
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			fmt.Printf("Error getting rows affected for job %d: %v\n", i, err)
			continue
		}

		if rowsAffected == 1 {
			inserted++
		} else if rowsAffected == 2 {
			updated++
		}
	}

	err = tx.Commit()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to commit transaction: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Bulk job creation completed. Inserted: %d, Updated: %d\n", inserted, updated)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{
		"inserted": inserted,
		"updated":  updated,
	})
}
