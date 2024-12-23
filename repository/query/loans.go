package query

const (
	CreateLoan = `
		INSERT INTO
			loans (
				borrower_id,
				principal_amount,
				rate,
				roi,
				agreement_letter
			)
			values (
				$1,
				$2,
				$3,
				$4,
				$5
			)
		RETURNING id;
	`

	GetLoanByID = `
		SELECT
			l.id,
			l.status,
			l.roi,
			l.principal_amount,
			COALESCE(SUM(m.funding_amount), 0) AS total_funding_amount,
			COALESCE(l.agreement_letter, '') AS agreement_letter
		FROM
			loans l
		LEFT JOIN
			marketplace m ON l.id = m.loan_id
		WHERE
			l.id = $1
		GROUP BY
			l.id, l.status, l.roi, l.principal_amount, l.agreement_letter;
	`

	ApproveLoan = `
		UPDATE
			loans
		SET
			status = 'approved',
			field_validator_id = $2,
			visit_doc_proof = $3,
			approval_date = $4
		WHERE
			id = $1;
	`

	FullyFundedLoan = `
		UPDATE
			loans
		SET
			status = 'invested'
		WHERE
			id = $1;
	`

	DisburseLoan = `
		UPDATE
			loans
		SET
			status = 'disbursed',
			agreement_letter = $2,
			disbursement_date = $3,
			collector_id = $4
		WHERE
			id = $1;
	`
)
