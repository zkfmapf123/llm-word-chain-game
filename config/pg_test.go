package config

import (
	"fmt"
	"testing"
)

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

func Test_QueryRow(t *testing.T) {
	pg := NewPGConn().WithHost(PG_HOST).WithPort(PG_PORT).WithUser(PG_USER).WithPassword(PG_PASSWORD).WithDBName(PG_DB_NAME)
	pg.MustConnect()

	var words []string
	rows, err := pg.DB.Query("select word from words where session_id=$1", 234932484)
	if err != nil {
		t.Error(err)
	}

	for rows.Next() {
		var word string
		rows.Scan(&word)
		words = append(words, word)
	}

	fmt.Println(words)
}

func Test_OverLabQuery(t *testing.T) {
	pg := NewPGConn().WithHost(PG_HOST).WithPort(PG_PORT).WithUser(PG_USER).WithPassword(PG_PASSWORD).WithDBName(PG_DB_NAME)
	pg.MustConnect()

	var exists bool
	err := pg.DB.QueryRow("select exists(select 1 from words where session_id=$1 and word=$2)", 1861074816, "장례").Scan(&exists)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(exists)
}
