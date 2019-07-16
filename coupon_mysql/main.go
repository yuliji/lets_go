package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
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

}
