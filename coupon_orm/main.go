package main

import (
	"fmt"
	"lets_go/coupon"
	"os"

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

func main() {

	o := orm.NewOrm()

	for i := 0; i < 200; i++ {
		coupon_code := coupon.GenerateCoupon()
		fmt.Println(coupon_code)
		coupon := Coupon{Coupon: coupon_code}
		id, err := o.Insert(&coupon)
		if err != nil {
			panic(err)
		}
		fmt.Printf("inserted id %d\n", id)

	}

}
