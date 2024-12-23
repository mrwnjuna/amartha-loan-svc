package dto

type Users struct {
	ID       int64  `json:"id" db:"id"`
	FullName string `json:"full_name" db:"full_name"`
	Email    string `json:"email" db:"email"`
	Type     string `json:"type" db:"type"`
}
