package repository

import (
	"amartha-loan-svc/dto"
	"amartha-loan-svc/repository/query"
)

//go:generate mockgen -destination=repository/mocks/mock_marketplace_repo.go -package=mocks amartha-loan-svc/repository MarketplaceRepo
type MarketplaceRepo interface {
	Fund(in dto.FundRequest) error
}

func (r AmarthaRepo) Fund(in dto.FundRequest) error {
	_, err := r.db.Exec(
		query.CreateFund,
		in.LoanID,
		in.LenderID,
		in.FundingAmount,
	)

	if err != nil {
		return err
	}

	return nil
}
