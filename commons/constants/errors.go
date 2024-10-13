package constants

const (
	PostgresConnectionError   = "failed to connect to Postgres: %w"
	ClosePostgresClientError  = "failed to close Postgres client"
	GetPostgresConfigError    = "failed to get Postgres config: %w"
	GetApplicationConfigError = "failed to get Application config: %w"
)

const (
	InvalidUsernameOrPasswordError = "invalid username or password"
	ProductNotFoundError           = "product not found"
)
