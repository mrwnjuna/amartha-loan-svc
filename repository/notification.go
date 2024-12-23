package repository

import (
	"amartha-loan-svc/dto"
	"fmt"
	"net/smtp"
	"os"
)

//go:generate mockgen -destination=repository/mocks/mock_notification_repo.go -package=mocks amartha-loan-svc/repository NotificationRepo
type NotificationRepo interface {
	SendEmail(in dto.SendEmailRequest) error
}

func (r AmarthaRepo) SendEmail(in dto.SendEmailRequest) error {

	from := os.Getenv("EMAIL_ADDRESS")
	password := os.Getenv("EMAIL_APP_PASSWORD") // Use the app-specific password here

	toEmailAddress := in.ToEmailAddress
	to := []string{toEmailAddress}

	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	address := host + ":" + port

	subject := fmt.Sprintf("Subject: Success Funding in Loan ID %d\n", in.LoanID)

	body := fmt.Sprintf(
		"This is to inform you that funding has been successfully made to Loan ID %d.\n\n"+
			"Loan Details:\n"+
			"- Loan ID: %d\n"+
			"- ROI: %.2f%%\n"+
			"- Agreement Letter: %s\n\n"+
			"Thank you for your participation in the loan funding process.",
		in.LoanID, in.LoanID, in.ROI, in.AgreementLetterURL,
	)

	message := []byte(subject + "\n\n" + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		return err
	}

	return nil
}
