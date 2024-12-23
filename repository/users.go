package repository

import (
	"amartha-loan-svc/dto"
	"amartha-loan-svc/repository/query"
)

//go:generate mockgen -destination=repository/mocks/mock_users_repo.go -package=mocks amartha-loan-svc/repository UsersRepo
type UsersRepo interface {
	GetUserByID(id int64) (out dto.Users, err error)
}

func (r AmarthaRepo) GetUserByID(id int64) (out dto.Users, err error) {
	err = r.db.QueryRow(
		query.GetUserByID,
		id,
	).Scan(
		&out.ID,
		&out.FullName,
		&out.Email,
		&out.Type,
	)

	if err != nil {
		return out, err
	}

	return out, nil
}
