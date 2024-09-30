package models

type Job struct {
	ID          int    `db:"id" json:"id,omitempty"`
	Title       string `db:"title" json:"title"`
	Company     string `db:"company" json:"company"`
	Location    string `db:"location" json:"location"`
	Salary      string `db:"salary" json:"salary"`
	Role        string `db:"role" json:"role"`
	Skills      string `db:"skills" json:"skills"`
	Remote      bool   `db:"remote" json:"remote"`
	Experience  string `db:"experience" json:"experience"`
	Education   string `db:"education" json:"education"`
	Department  string `db:"department" json:"department"`
	JobType     string `db:"job_type" json:"jobType"`
	URL         string `db:"url" json:"url"`
	Description string `db:"description" json:"description"`
}
