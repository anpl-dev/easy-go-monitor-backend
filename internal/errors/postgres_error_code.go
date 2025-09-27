package errors

// Postgres error codes (https://www.postgresql.jp/document/15/html/errcodes-appendix.html)
const (
	ForeignKeyViolation = "23503"
	UniqueViolation     = "23505"
	NotNullViolation    = "23502"
	CheckViolation      = "23514"
)
