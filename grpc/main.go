package main

import fmt "fmt"

func main() {
	p := &Person{
		Name: "Levi",
	}

	fmt.Println(p.Descriptor())
}
