package repository

import (
	"amartha-loan-svc/dto"
	"amartha-loan-svc/repository/query"
	"time"
)

//go:generate mockgen -destination=repository/mocks/mock_loans_repo.go -package=mocks amartha-loan-svc/repository LoansRepo
type LoansRepo interface {
	CreateLoan(in dto.CreateLoan) (out dto.Loans, err error)
	GetLoanByID(id int64) (out dto.Loans, err error)
	ApproveLoan(in dto.ApproveLoan) error
	FullyFundedLoan(id int64) error
	DisburseLoan(in dto.DisburseLoan) error
}

func (r AmarthaRepo) CreateLoan(in dto.CreateLoan) (out dto.Loans, err error) {
	err = r.db.QueryRow(
		query.CreateLoan,
		in.BorrowerID,
		in.PrincipalAmount,
		in.Rate,
		in.ROI,
		in.AgreementLetter,
	).Scan(
		&out.ID,
	)

	if err != nil {
		return dto.Loans{}, err
	}

	return out, nil
}

func (r AmarthaRepo) GetLoanByID(id int64) (out dto.Loans, err error) {
	err = r.db.QueryRow(
		query.GetLoanByID,
		id,
	).Scan(
		&out.ID,
		&out.Status,
		&out.ROI,
		&out.PrincipalAmount,
		&out.TotalFundingAmount,
		&out.AgreementLetter,
	)

	if err != nil {
		return out, err
	}

	return out, nil
}

func (r AmarthaRepo) ApproveLoan(in dto.ApproveLoan) error {
	_, err := r.db.Exec(
		query.ApproveLoan,
		in.LoanID,
		in.FieldValidatorID,
		in.VisitDocProof,
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

func (r AmarthaRepo) FullyFundedLoan(id int64) error {
	_, err := r.db.Exec(
		query.FullyFundedLoan,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r AmarthaRepo) DisburseLoan(in dto.DisburseLoan) error {
	_, err := r.db.Exec(
		query.DisburseLoan,
		in.LoanID,
		in.SignedAgreementLetter,
		time.Now(),
		in.CollectorID,
	)

	if err != nil {
		return err
	}

	return nil
}
