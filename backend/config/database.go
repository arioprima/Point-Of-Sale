package config

import (
	"database/sql"
	"time"
)

func ConnectionDB() *sql.DB {
	//mysql://ubuntu:TanahDamai@123@tcp(192.168.221.97:3306)/belajar_golang_jwt
	db, err := sql.Open("mysql", "centos:TanahDamai@123@tcp(103.127.97.105:3306)/belajar_golang_jwt")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(25)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
