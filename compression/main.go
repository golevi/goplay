//
// compression
//
// references used to write this
// - https://en.wikipedia.org/wiki/Huffman_coding
// - https://golang.org/ref/spec
// - https://golang.org/pkg/bytes/
//
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

var input string

func main() {
	var symbols = make(map[byte]int)
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, i := range input {
		_, ok := symbols[i]
		if ok {
			symbols[i]++
		} else {
			symbols[i] = 1
		}
	}
	var symbolsCount = make(map[int]byte)

	for k, v := range symbols {
		symbolsCount[v] = k
	}
	bit := 1
	f := 8
	fmt.Println(symbolsCount)
	fmt.Printf("%08b %08b %08b\n", ((bit|f)<<4)|f, bit, f)
	fmt.Println(((bit | f) << 4) | f)
	fmt.Printf("%08b", 1|4)
}
