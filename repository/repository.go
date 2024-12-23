package repository

import "database/sql"

type AmarthaRepo struct {
	db *sql.DB
}

type AmarthaRepoInterface interface {
	LoansRepo
	UsersRepo
	MarketplaceRepo
	NotificationRepo
}

func NewAmarthaRepo(db *sql.DB) *AmarthaRepo {
	return &AmarthaRepo{db: db}
}
