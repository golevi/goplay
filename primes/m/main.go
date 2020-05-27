package main

import (
	"context"
	"fmt"
)

type prime struct {
	P string
	N int64
}

func primeSearcher(ctx context.Context) {

}

func main() {
	fmt.Println("Starting")

	ctx := context.TODO()

	for i := 0; i < 10; i++ {
		go primeSearcher(ctx)
	}
}
