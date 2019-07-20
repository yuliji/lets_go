package coupon

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func GenerateCoupon() string {
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

	var buffer strings.Builder

	charset_len := len(charset)
	var max big.Int
	max.SetInt64(int64(charset_len))

	for i := 0; i < 16; i++ {
		randInt, err := rand.Int(rand.Reader, &max)
		if err != nil {
			panic(err)
		}
		idx := int(randInt.Int64())
		buffer.WriteString(charset[idx])
		if i == 3 || i == 7 || i == 11 {
			buffer.WriteString("-")
		}
	}
	coupon := buffer.String()
	return coupon
}
