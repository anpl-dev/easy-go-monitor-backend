package codes

// Postgres error codes (https://www.postgresql.jp/document/15/html/errcodes-appendix.html)
const (
	PostgresForeignKeyViolation = "23503"
	PostgresUniqueViolation     = "23505"
	PostgresNotNullViolation    = "23502"
	PostgresCheckViolation      = "23514"
)
