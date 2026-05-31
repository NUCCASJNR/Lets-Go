package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func initDB() {
	var err error
	DB, err = sql.Open(
	"mysql",
	"course_user:Course_pwd123@@tcp(127.0.0.1:3306)/course_db",
)
	
	if err != nil {
		fmt.Println("DB connection error:", err)
		return
	}

	err = DB.Ping()
	if err != nil {
		fmt.Println("DB not reachable:", err)
		return
	}

	fmt.Println("Connected to database ✅")
}