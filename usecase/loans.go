package usecase

import (
	"amartha-loan-svc/dto"
	"amartha-loan-svc/utils"
	"database/sql"
	"fmt"
)

type LoansUsecase interface {
	CreateLoan(in dto.CreateLoan) (out dto.Loans, appErr *utils.AppError)
	ApproveLoan(in dto.ApproveLoan) *utils.AppError
	DisburseLoan(in dto.DisburseLoan) *utils.AppError
}

func (u *AmarthaUsecase) CreateLoan(in dto.CreateLoan) (out dto.Loans, appErr *utils.AppError) {
	user, err := u.amarthaRepo.GetUserByID(in.BorrowerID)
	if err == sql.ErrNoRows {
		return dto.Loans{}, utils.NewNotFoundError("borrower not found", fmt.Sprintf("borrower id: %d", in.BorrowerID))
	}
	if err != nil {
		return dto.Loans{}, utils.NewInternalServerError("failed to retrieve loan", err.Error())
	}

	if user.Type != "borrower" {
		return dto.Loans{}, utils.NewBadRequestError("user can't create loan", fmt.Sprintf("user type: %s", user.Type))
	}

	out, err = u.amarthaRepo.CreateLoan(in)
	if err != nil {
		return dto.Loans{}, utils.NewInternalServerError("failed to create loan", err.Error())
	}

	return out, nil
}

func (u *AmarthaUsecase) ApproveLoan(in dto.ApproveLoan) *utils.AppError {
	loan, err := u.amarthaRepo.GetLoanByID(in.LoanID)
	if err == sql.ErrNoRows {
		return utils.NewNotFoundError("loan not found", fmt.Sprintf("loan id: %d", in.LoanID))
	}
	if err != nil {
		return utils.NewInternalServerError("failed to retrieve loan", err.Error())
	}

	if loan.Status != "proposed" {
		return utils.NewBadRequestError("loan is not proposed", fmt.Sprintf("loan status: %s", loan.Status))
	}

	err = u.amarthaRepo.ApproveLoan(in)
	if err != nil {
		return utils.NewInternalServerError("failed to approve loan", err.Error())
	}

	return nil
}

func (u *AmarthaUsecase) DisburseLoan(in dto.DisburseLoan) *utils.AppError {
	loan, err := u.amarthaRepo.GetLoanByID(in.LoanID)
	if err == sql.ErrNoRows {
		return utils.NewNotFoundError("loan not found", fmt.Sprintf("loan id: %d", in.LoanID))
	}
	if err != nil {
		return utils.NewInternalServerError("failed to retrieve loan", err.Error())
	}

	if loan.Status != "invested" {
		return utils.NewBadRequestError("loan is not invested", fmt.Sprintf("loan status: %s", loan.Status))
	}

	err = u.amarthaRepo.DisburseLoan(in)
	if err != nil {
		return utils.NewInternalServerError("failed to disburse loan", err.Error())
	}

	return nil
}
