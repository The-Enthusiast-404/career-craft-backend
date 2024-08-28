package models

type Job struct {
	ID      int    `db:"id" json:"id,omitempty"`
	Title   string `db:"title" json:"title"`
	Company string `db:"company" json:"company"`
	URL     string `db:"url" json:"url"`
}
