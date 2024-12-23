package usecase

import (
	"amartha-loan-svc/dto"
	"amartha-loan-svc/utils"
	"database/sql"
	"fmt"
)

type MarketplaceUsecase interface {
	Fund(in dto.FundRequest) *utils.AppError
}

func (u *AmarthaUsecase) Fund(in dto.FundRequest) *utils.AppError {
	//TODO: use trx
	user, err := u.amarthaRepo.GetUserByID(in.LenderID)
	if err == sql.ErrNoRows {
		return utils.NewNotFoundError("lender not found", fmt.Sprintf("lender id: %d", in.LenderID))
	}
	if err != nil {
		return utils.NewInternalServerError("failed to retrieve user", err.Error())
	}

	if user.Type != "lender" {
		return utils.NewBadRequestError("user can't fund", fmt.Sprintf("user type: %s", user.Type))
	}

	loan, err := u.amarthaRepo.GetLoanByID(in.LoanID)
	if err == sql.ErrNoRows {
		return utils.NewNotFoundError("loan not found", fmt.Sprintf("loan id: %d", in.LoanID))
	}
	if err != nil {
		return utils.NewInternalServerError("failed to retrieve loan", err.Error())
	}

	if loan.Status != "approved" {
		return utils.NewBadRequestError("loan is not approved", fmt.Sprintf("loan status: %s", loan.Status))
	}

	remaining := loan.PrincipalAmount - loan.TotalFundingAmount
	if in.FundingAmount > remaining {
		return utils.NewBadRequestError("exceeds funding amount", fmt.Sprintf("remaining amount: %0.3f", remaining))
	}

	err = u.amarthaRepo.Fund(in)
	if err != nil {
		return utils.NewInternalServerError("failed to fund loan", err.Error())
	}

	loan, err = u.amarthaRepo.GetLoanByID(in.LoanID)
	if err != nil {
		return utils.NewInternalServerError("failed to retrieve loan", err.Error())
	}

	if loan.TotalFundingAmount == loan.PrincipalAmount {
		err = u.amarthaRepo.FullyFundedLoan(loan.ID)
		if err != nil {
			return utils.NewInternalServerError("failed to set fully funded loan", err.Error())
		}

		err = u.amarthaRepo.SendEmail(dto.SendEmailRequest{
			ToEmailAddress:     user.Email,
			LoanID:             loan.ID,
			ROI:                loan.ROI,
			AgreementLetterURL: loan.AgreementLetter,
		})
		if err != nil {
			return utils.NewInternalServerError("failed to send email", err.Error())
		}
	}

	return nil
}
