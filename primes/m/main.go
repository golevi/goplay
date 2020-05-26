package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var knowns = []uint64{
	82589933,
}

func mersenne(n int64) {
	b, y, m := &big.Int{}, &big.Int{}, &big.Int{}

	y.SetInt64(n)
	m.SetInt64(0)
	b.SetInt64(2)

	b.Exp(b, y, m)

	x := &big.Int{}
	x.SetInt64(1)
	b.Sub(b, x)

	if b.ProbablyPrime(16) {
		fmt.Printf("\n================\n2^n-1 | n=%d\n================\n", n)
		fmt.Println(b.String())
	}
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 1 {
		fmt.Println("please provide a starting number")
		return
	}

	s, err := strconv.Atoi(argsWithoutProg[0])
	if err != nil {
		panic(err)
	}
	x := int64(s)

	for i := x; i < x*10; i++ {
		mersenne(i)
	}
}
