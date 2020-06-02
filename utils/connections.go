package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Connecttodb() (connection *sql.DB, err error) {
	dburi := "root:daddy@tcp(localhost:3306)/singers"
	connection, err = sql.Open("mysql", dburi)
	if err != nil {
		fmt.Println("unable to connect to db")
		return
	}
	return
}
