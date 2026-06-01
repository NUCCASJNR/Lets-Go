package config

import (
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error

	DB, err = sql.Open("mysql", "m_user:M_pwd123@tcp(127.0.0.1:3306)/movies_db")
	if err != nil {
		fmt.Println("DB error:", err)
		return
	}

	err = DB.Ping()
	if err != nil {
		fmt.Println("DB not reachable:", err)
		return
	}

	fmt.Println("Database connected ✅")
}