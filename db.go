package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB
var logined bool
var FirstFreeID int = 0

func cfg() {
	cfg := mysql.NewConfig()
	(*cfg).User = "root"
	(*cfg).Addr = "localhost"
	(*cfg).Passwd = "nikita2005"
	(*cfg).Net = "tcp"
	(*cfg).DBName = "chat"
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println(err)
		return
	}
}
