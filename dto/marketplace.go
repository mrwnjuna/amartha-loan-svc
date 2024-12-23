package dto

type (
	FundRequest struct {
		LoanID        int64   `json:"loan_id" validate:"required"`
		LenderID      int64   `json:"lender_id" validate:"required"`
		FundingAmount float64 `json:"funding_amount" validate:"required"`
	}
)
