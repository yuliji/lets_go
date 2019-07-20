package main

import (
	"database/sql"
	"fmt"
	"lets_go/coupon"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db_username := os.Getenv("DB_USERNAME")
	db_passwd := os.Getenv("DB_PASSWORD")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/coupons", db_username, db_passwd)

	db, err := sql.Open("mysql", dataSourceName)
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
		coupon := coupon.GenerateCoupon()
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
