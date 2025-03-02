package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type PGConn struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	DB       *sql.DB
}

func NewPGConn() *PGConn {
	return &PGConn{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
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

func (pg *PGConn) MustConnect() *PGConn {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", pg.Host, pg.Port, pg.User, pg.Password, pg.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	pg.DB = db
	return pg
}

func (pg *PGConn) Close() {
	pg.DB.Close()
}
