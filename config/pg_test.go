package config

import "testing"

const (
	PG_HOST     = "localhost"
	PG_PORT     = "5432"
	PG_USER     = "postgres"
	PG_PASSWORD = "postgres"
	PG_DB_NAME  = "vectordb"
)

func Test_PgConnection(t *testing.T) {
	pg := NewPGConn().WithHost(PG_HOST).WithPort(PG_PORT).WithUser(PG_USER).WithPassword(PG_PASSWORD).WithDBName(PG_DB_NAME)
	pg.MustConnect()
}
