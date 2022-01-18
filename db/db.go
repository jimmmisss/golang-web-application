package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Conexao() *sql.DB {
	conexao := "user=lezzr_go dbname=lezzr_go password=lezzr_go port=5431 host=127.0.0.1 sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
