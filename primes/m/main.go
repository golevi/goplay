package main

import (
	"fmt"
	"math/big"
)

var knowns = []uint64{
	82589933,
}

func mersenne(p int64) {
	b, y, m := &big.Int{}, &big.Int{}, &big.Int{}

	y.SetInt64(p)
	m.SetInt64(0)
	b.SetInt64(2)

	b.Exp(b, y, m)

	x := &big.Int{}
	x.SetInt64(1)
	b.Sub(b, x)

	if b.ProbablyPrime(16) {
		fmt.Printf("p: %d\n\n", p)
		fmt.Println(b.String())
	}
}

func main() {
	mersenne(1279)
}
