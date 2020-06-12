package main

import (
	"errors"
	"fmt"
)

// Bit represents a binary bit
type Bit bool

// BinaryNumber represents a 4 bit binary number
type BinaryNumber struct {
	Bit1 Bit
	Bit2 Bit
	Bit3 Bit
	Bit4 Bit
	Bit5 Bit
}

// New BinaryNumber
func New(num int) (*BinaryNumber, error) {
	bn := &BinaryNumber{}
	if num > 8 {
		return bn, errors.New("only 4 bits")
	}

	if num == 8 {
		bn.Bit1 = true
		bn.Bit2 = true
		bn.Bit3 = true
		bn.Bit4 = true

		return bn, nil
	}

	if num == 7 {
		bn.Bit1 = true
		bn.Bit2 = true
		bn.Bit3 = true

		return bn, nil
	}

	if num == 6 {
		bn.Bit2 = true
		bn.Bit3 = true

		return bn, nil
	}

	if num == 5 {
		bn.Bit1 = true
		bn.Bit4 = true

		return bn, nil
	}

	if num == 4 {
		bn.Bit3 = true

		return bn, nil
	}

	if num == 3 {
		bn.Bit1 = true
		bn.Bit2 = true

		return bn, nil
	}

	if num == 2 {
		bn.Bit2 = true

		return bn, nil
	}

	if num == 1 {
		bn.Bit1 = true

		return bn, nil
	}

	return bn, nil
}

// Dec converts the BinaryNumber to an int
func (b BinaryNumber) Dec() int {
	total := 0
	if b.Bit1 {
		total = total + 1
	}

	if b.Bit2 {
		total = total + 2
	}

	if b.Bit3 {
		total = total + 4
	}

	if b.Bit4 {
		total = total + 8
	}

	return total
}

// Gate is a gate
type Gate func(a, b Bit) Bit

func xor(a, b Bit) Bit {
	return a != b
}

func and(a, b Bit) Bit {
	return a == true && b == true
}

func or(a, b Bit) Bit {
	return a || b
}

func add(a, b, c Bit) (Bit, Bit) {
	var sum, carry Bit

	AxB := xor(a, b)
	AxBxC := xor(AxB, c)

	AaB := and(a, b)
	AxBaC := and(AxB, c)

	AxBaCoAaB := or(AxBaC, AaB)

	sum = AxBxC
	carry = AxBaCoAaB

	return sum, carry
}

func addition(x, y *BinaryNumber) {
	b1, o := add(x.Bit1, y.Bit1, false)
	b2, o := add(x.Bit2, y.Bit2, o)
	b3, o := add(x.Bit3, y.Bit3, o)
	b4, o := add(x.Bit4, y.Bit4, o)

	bn := &BinaryNumber{
		Bit1: b1,
		Bit2: b2,
		Bit3: b3,
		Bit4: b4,
		Bit5: o,
	}

	fmt.Println(bn)
	fmt.Println(bn.Dec())
}

func main() {

	seven, _ := New(7)
	fmt.Println(seven.Dec())

	three, _ := New(4)
	fmt.Println(three.Dec())

	addition(seven, three)
}
