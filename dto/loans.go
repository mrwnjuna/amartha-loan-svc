package dto

import "time"

type (
	CreateLoan struct {
		BorrowerID      int64   `json:"borrower_id" validate:"required"`
		PrincipalAmount float64 `json:"principal_amount" validate:"required"`
		Rate            float32 `json:"rate" validate:"required"`
		ROI             float32 `json:"roi" validate:"required"`
		AgreementLetter string  `json:"agreement_letter" validate:"required"`
	}

	ApproveLoan struct {
		LoanID           int64  `json:"loan_id" validate:"-"`
		VisitDocProof    string `json:"visit_doc_proof" validate:"required"`
		FieldValidatorID string `json:"field_validator_id" validate:"required"`
	}

	DisburseLoan struct {
		LoanID                int64  `json:"loan_id" validate:"required"`
		SignedAgreementLetter string `json:"signed_agreement_letter" validate:"required"`
		CollectorID           string `json:"collector_id" validate:"required"`
	}

	Loans struct {
		ID                 int64   `json:"id" db:"id"`
		BorrowerID         int64   `json:"borrower_id" db:"borrower_id"`
		PrincipalAmount    float64 `json:"principal_amount" db:"principal_amount"`
		TotalFundingAmount float64 `json:"total_funding_amount" db:"total_funding_amount"`

		Rate             float32   `json:"rate" db:"rate"`
		ROI              float32   `json:"roi" db:"roi"`
		AgreementLetter  string    `json:"agreement_letter" db:"agreement_letter"`
		Status           string    `json:"status" db:"status"`
		FieldValidatorID string    `json:"field_validator_id" db:"field_validator_id"`
		VisitDocProof    string    `json:"visit_doc_proof" db:"visit_doc_proof"`
		ApprovalDate     time.Time `json:"approval_date" db:"approval_date"`
		CollectorID      string    `json:"collector_id" db:"collector_id"`
		DisbursementDate time.Time `json:"disbursement_date" db:"disbursement_date"`
	}
)
