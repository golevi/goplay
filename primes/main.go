package main

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

const max string = "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"

func findPrime(ctx context.Context, i chan<- string) {
	b := &big.Int{}
	n := &big.Int{}

	n.SetString(max, 16)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			p := b.Rand(r, n)

			if p.ProbablyPrime(4096) {
				i <- p.String()
			}
		}
	}
}

func printPrime(ctx context.Context, i <-chan string) {
	for {
		select {
		case p := <-i:
			fmt.Println(p)
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	fmt.Println("Starting")

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	defer cancel()

	ch := make(chan string)

	for i := 0; i < 10; i++ {
		go findPrime(ctx, ch)
	}

	go printPrime(ctx, ch)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Quit")
			return
		}
	}
}
