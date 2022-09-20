package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host   = "127.0.0.1"
	port   = "3306"
	user   = "root"
	pass   = ""
	dbname = "hacktiv8-mnc"
)

var DB *sql.DB

func ConnectPostgres() error {
	dsn := fmt.Sprintf("%s:%s@/%s", user, pass, dbname)
	// dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)

	log.Println(dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return err
	}

	DB = db

	return nil
}

func GetPostgres() *sql.DB {
	return DB
}
