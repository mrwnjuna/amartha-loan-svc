package dto

type (
	SendEmailRequest struct {
		ToEmailAddress     string
		LoanID             int64
		ROI                float32
		AgreementLetterURL string
	}
)
