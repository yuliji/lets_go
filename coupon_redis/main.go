package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/go-redis/redis"

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
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	coupons_set := "coupons"

	for i := 0; i < 200; i++ {
		coupon_code := generate_coupon()
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
