package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() { 
	var err error
	db, err = sql.Open("mysql", "root:ML$P-%J%+2-9z$Y&@/securitydb")
	if err != nil {
	panic(err)
	} 
	defer db.Close()
}