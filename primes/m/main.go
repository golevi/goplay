package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var knowns = []uint64{
	82589933,
}

type prime struct {
	P string
	N int64
}

func mersenne(ctx context.Context, pch chan prime, n int64) {
	b, y, m := &big.Int{}, &big.Int{}, &big.Int{}

	y.SetInt64(n)
	m.SetInt64(0)
	b.SetInt64(2)

	// first check if n is prime, if it isn't, return
	if !y.ProbablyPrime(4) {
		return
	}

	// Do this to show progress
	fmt.Printf(".")

	b.Exp(b, y, m)

	x := &big.Int{}
	x.SetInt64(1)
	b.Sub(b, x)

	if b.ProbablyPrime(16) {
		var p prime
		p.P = b.String()
		p.N = n

		pch <- p
	}
}

func printer(ctx context.Context, pch chan prime, cancel context.CancelFunc) {
	for {
		select {
		case p := <-pch:
			fmt.Printf("\n================\n%d (2^n)-1\n================\n%s\n\n", p.N, p.P)
			cancel()
		case <-ctx.Done():
			return
		}
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

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	pch := make(chan prime)

	x := int64(s)
	for i := x; i < x*10; i = i + 2 {
		go mersenne(ctx, pch, i)
		go mersenne(ctx, pch, i+1)
	}

	printer(ctx, pch, cancel)
}
