package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
)

type Coupon struct {
	Id     int    `orm:"auto"`
	Coupon string `orm:"size(19)"`
}

func init() {
	// register model
	orm.RegisterModel(new(Coupon))

	db_username := os.Getenv("DB_USERNAME")
	db_passwd := os.Getenv("DB_PASSWORD")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/coupons2?charset=utf8",
		db_username, db_passwd)

	// set default database
	orm.RegisterDataBase("default", "mysql", dataSourceName, 30)

	// create table
	orm.RunSyncdb("default", false, true)
}

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

	o := orm.NewOrm()

	for i := 0; i < 200; i++ {
		coupon_code := generate_coupon()
		fmt.Println(coupon_code)
		coupon := Coupon{Coupon: coupon_code}
		id, err := o.Insert(&coupon)
		if err != nil {
			panic(err)
		}
		fmt.Printf("inserted id %d\n", id)

	}

}
