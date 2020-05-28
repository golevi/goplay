package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"time"
)

const max string = "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"

func findPrime(ctx context.Context, i chan<- string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			b, err := rand.Prime(rand.Reader, 4096)
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
