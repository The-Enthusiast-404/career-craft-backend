package models

type Company struct {
	ID          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	Website     string `db:"website" json:"website"`
	Industry    string `db:"industry" json:"industry"`
	Size        string `db:"size" json:"size"`
	Founded     int    `db:"founded" json:"founded"`
}
