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
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3306)/coupons")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	stmt, err := db.Prepare("INSERT INTO coupons(coupon) VALUES(?)")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 200; i++ {
		coupon := generate_coupon()
		fmt.Println(coupon)

		res, err := stmt.Exec(coupon)
		if err != nil {
			panic(err)
		}
		lastId, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			panic(err)
		}
		fmt.Println("ID = %d, affected = %d\n", lastId, rowCnt)
	}

}
