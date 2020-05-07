//
// rsa
//
// references used to write this
// - https://blog.cloudflare.com/a-relatively-easy-to-understand-primer-on-elliptic-curve-cryptography/
// - https://en.wikipedia.org/wiki/RSA_(cryptosystem)
// - https://golang.org/pkg/math/big/
//
package main

import (
	"fmt"
	"math/big"
)

func main() {
	var plainText = "ACLOUD"

	// 1. Choose two distinct prime numbers, p and q
	// 4. Choose any number 1 < e < 780 that is coprime to 780.
	var p, q, e big.Int
	p.SetInt64(61)
	q.SetInt64(53)
	e.SetInt64(17)

	// 2. Compute n = pq.
	// n is used as the modulus for both the public and private keys. Its length, usually expressed in bits, is the key length.
	// n is released as part of the public key.
	var n big.Int
	n = *n.Mul(&p, &q)

	// 3. Compute the Carmichael's totient function of the product as λ(n) = lcm(p − 1, q − 1) giving
	var lcm, lcmp, lcmq, i big.Int
	i.SetInt64(1)
	lcmp.Sub(&p, &i)
	lcmq.Sub(&q, &i)
	lcm.Mul(lcm.Div(&lcmp, lcm.GCD(nil, nil, &lcmp, &lcmq)), &lcmq)

	// 5. Compute d, the modular multiplicative inverse of e (mod λ(n)) yielding,
	var d big.Int
	d.ModInverse(&e, &lcm)

	fmt.Println("Input")
	fmt.Println("=====")
	fmt.Printf("%s\n\n", plainText)
	fmt.Println("Parameters")
	fmt.Println("==========")
	fmt.Printf("p:\t%s\n", p.String())
	fmt.Printf("q:\t%s\n", q.String())
	fmt.Printf("e:\t%s\n", e.String())
	fmt.Printf("n:\t%s\n", n.String())
	fmt.Printf("lcm:\t%s\n", lcm.String())
	fmt.Printf("inv:\t%s\n\n", d.String())

	// Encrypt
	fmt.Println("input\toutput")
	fmt.Println("-----\t------")
	var l, t big.Int
	var ciphertext []string
	for _, c := range plainText {
		l.SetInt64(int64(c))
		t.SetInt64(int64(c))
		t.Exp(&l, &e, &n)
		fmt.Printf("%s\t%s\n", l.String(), t.String())
		ciphertext = append(ciphertext, t.String())
	}
	fmt.Printf("\nCiphertext: %s\n\n", ciphertext)

	fmt.Println("input\toutput")
	fmt.Println("-----\t------")
	var decode []byte
	// Decrypt
	for _, k := range ciphertext {
		l.SetString(k, 10)
		t.Exp(&l, &d, &n)
		fmt.Printf("%s\t%s\n", l.String(), t.String())
		decode = append(decode, byte(t.Int64()))
	}
	fmt.Printf("\nDecrypted: %s\n", string(decode))
}
