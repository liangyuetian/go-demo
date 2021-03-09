package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const (
	userName = "root"
	password = "123456"
	ip       = "localhost"
	port     = "3306"
	dbName   = "l"
)

var DB *sql.DB

func InitDB() {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	DB, _ = sql.Open("mysql", path)
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)

	if err := DB.Ping(); err != nil {
		fmt.Println("msql open database fail")
		return
	}
	fmt.Println("mysql connect success")
}

func GetDB() *sql.DB {
	return DB
}
