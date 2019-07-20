package main

import (
	"fmt"
	"lets_go/coupon"

	"github.com/go-redis/redis"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	coupons_set := "coupons"

	for i := 0; i < 200; i++ {
		coupon_code := coupon.GenerateCoupon()
		fmt.Println(coupon_code)
		err = client.SAdd(coupons_set, coupon_code).Err()
		if err != nil {
			panic(err)
		}
	}
	val, err := client.SMembers(coupons_set).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("coupons", val)

}
