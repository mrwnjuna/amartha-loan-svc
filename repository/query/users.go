package query

const (
	GetUserByID = `
		SELECT
			id,
			full_name,
			email,
			type
		FROM
			users
		WHERE
			id = $1;
	`
)
