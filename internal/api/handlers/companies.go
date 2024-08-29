package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"dev.theenthusiast.career-craft/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
)

type CompanyHandler struct {
	DB *sqlx.DB
}

func NewCompanyHandler(db *sqlx.DB) *CompanyHandler {
	return &CompanyHandler{DB: db}
}

func (h *CompanyHandler) GetCompanyDetails(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	companyName := ps.ByName("company")

	var company models.Company
	err := h.DB.Get(&company, "SELECT * FROM companies WHERE name = $1", companyName)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Company not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(company)
}
