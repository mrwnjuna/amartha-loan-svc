package query

const (
	CreateFund = `
		INSERT INTO
			marketplace (
				loan_id,
				lender_id,
				funding_amount
			)
			values (
				$1,
				$2,
				$3
			);
	`
)
