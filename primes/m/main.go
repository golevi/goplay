package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

const primeSize = 27 // ~111,926,261

func mersenne(ctx context.Context, i <-chan string, cancel context.CancelFunc) {
	for {
		select {
		case p := <-i:
			//
			b, y, m := &big.Int{}, &big.Int{}, &big.Int{}

			y.SetString(p, 10)
			m.SetInt64(0)
			b.SetInt64(2)

			// Do this to show progress
			log.Println(".")

			b.Exp(b, y, m)

			x := &big.Int{}
			x.SetInt64(1)
			b.Sub(b, x)

			// fmt.Println(y.String())
			// fmt.Println(b.String())

			if b.ProbablyPrime(16) {
				fmt.Println(b.String())
				fmt.Println(p)
				cancel()
			}
			//
		case <-ctx.Done():
			return
		}
	}
}

func findPrime(ctx context.Context, i chan<- string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			b, err := rand.Prime(rand.Reader, primeSize)
			if err != nil {
				log.Println(err)
			}
			i <- b.String()
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

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	ch := make(chan string)

	// go printPrime(ctx, ch)
	go findPrime(ctx, ch)

	for i := 0; i < 10; i++ {
		go mersenne(ctx, ch, cancel)
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Quit")
			return
		}
	}
}
