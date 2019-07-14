package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func main() {
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
	count := 0

	coupons := make(map[string]bool)

	for count < 200 {
		var buffer bytes.Buffer

		for i := 0; i < 16; i++ {
			buffer.WriteString(charset[rand.Intn(len(charset))])
			if i == 3 || i == 7 || i == 11 {
				buffer.WriteString("-")
			}
		}
		coupon := buffer.String()
		_, present := coupons[coupon]
		if !present {
			fmt.Println(coupon)
			coupons[coupon] = true
			count++
		}
	}

}
