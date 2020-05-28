package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"
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
			log.Printf("Checking %s\n", p)

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
				return
			}
			log.Printf("Negative %s\n", p)
			//
		case <-ctx.Done():
			return
		}
	}
}

func findPrime(ctx context.Context, cancel context.CancelFunc, i chan<- string, size int) {
	var checked map[string]bool
	checked = make(map[string]bool)
	iteration := 0
	for {
		select {
		case <-ctx.Done():
			return
		default:
			iteration++
			b, err := rand.Prime(rand.Reader, size)
			if err != nil {
				log.Println(err)
			}
			if checked[b.String()] {
				log.Printf("Skipping %s\n", b.String())
				// some exponents were only generating a handful of primes, so it
				// would try the same number over and over, forever. now it checks
				// to see how many times it has checked the same number. if it is
				// a lot, it quits.
				if iteration > 10000 {
					cancel()
				}
				continue
			}
			i <- b.String()

			checked[b.String()] = true
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

func ping(ctx context.Context, sleep int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Print(".")
			time.Sleep(time.Second * sleep)
		}
	}
}

func main() {
	fmt.Println("Starting")

	args := os.Args[1:]
	size, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	ch := make(chan string)

	// go printPrime(ctx, ch)
	go findPrime(ctx, cancel, ch, size)
	go ping(ctx, 10)

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
