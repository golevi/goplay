//
// XOR example
//
package main

import "fmt"

func main() {
	var a int = 22
	var x int = 99
	var b, c int

	fmt.Printf("a: %08b\nx: %08b\n\n", a, x)

	b = a ^ x
	fmt.Println("b = a ^ x")
	fmt.Printf("b: %08b\n\n", b)

	c = b ^ x
	fmt.Println("c = b ^ x")
	fmt.Printf("c: %08b\n\n", c)

	fmt.Println("Does a == c?")
	if a == c {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
