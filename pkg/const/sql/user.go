package sql

const (
	CreateUserQuery = `INSERT INTO users (user_name, email, hashed_password, created_at, updated_at)
											VALUES ($1, $2, $3, now(), now()) 
						          RETURNING *`
	)
