package models

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	USERNAME = "root"
	PASSWORD = "12345678"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "blog"
)

func connect() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return DB, err
	}
	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return DB, err
}
func queryOne(DB *sql.DB) {
	fmt.Println("query times:", 1)
	user := new(User)
	row := DB.QueryRow("select * from users where id=?", 1)
	if err := row.Scan(&user.id, &user.name, &user.email); err != nil {
		fmt.Printf("scan failed, err:%v", err)
		return
	}
	fmt.Println(*user)
}

func init() {
	DB, err := connect()
	if err != nil {
		fmt.Println("db connect error:%v", err)
	}
	queryOne(DB)
}
