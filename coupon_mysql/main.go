package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func generate_coupon() string {
	// We do not use 1 I 0 o
	charset := []string{
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
		"G",
		"H",
		"J",
		"K",
		"L",
		"N",
		"N",
		"P",
		"Q",
		"R",
		"S",
		"T",
		"U",
		"V",
		"W",
		"X",
		"Y",
		"Z",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
	}

	rand.Seed(time.Now().UnixNano())

	var buffer strings.Builder

	for i := 0; i < 16; i++ {
		buffer.WriteString(charset[rand.Intn(len(charset))])
		if i == 3 || i == 7 || i == 11 {
			buffer.WriteString("-")
		}
	}
	coupon := buffer.String()
	return coupon
}

func main() {
	for i := 0; i < 200; i++ {
		coupon := generate_coupon()
		fmt.Println(coupon)
	}

	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	hasDb := false
	databases := showDatabases(db)
	for _, database := range databases {
		if database == "coupons" {
			hasDb = true
		}
	}

	if !hasDb {
		createDatabase(db)
		fmt.Println("created db")
	} else {
		fmt.Println("db already exists")
	}

}

func createDatabase(db *sql.DB) {
	stmt, err := db.Prepare("CREATE DATABASE coupons")
	if err != nil {
		panic(err)
	}
	_, err2 := stmt.Exec()
	if err2 != nil {
		panic(err)
	}
}

func showDatabases(db *sql.DB) []string {
	dataBases := make([]string, 10)

	var database string
	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&database)
		if err != nil {
			panic(err.Error())
		}
		dataBases = append(dataBases, database)
	}
	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
	return dataBases
}
