package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PGConn struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewPGConn() *PGConn {
	return &PGConn{}
}

func (pg *PGConn) WithHost(host string) *PGConn {
	pg.Host = host
	return pg
}

func (pg *PGConn) WithPort(port string) *PGConn {
	pg.Port = port
	return pg
}

func (pg *PGConn) WithUser(user string) *PGConn {
	pg.User = user
	return pg
}

func (pg *PGConn) WithPassword(password string) *PGConn {
	pg.Password = password
	return pg
}

func (pg *PGConn) WithDBName(dbName string) *PGConn {
	pg.DBName = dbName
	return pg
}

func (pg *PGConn) MustConnect() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", pg.Host, pg.Port, pg.User, pg.Password, pg.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()
}
