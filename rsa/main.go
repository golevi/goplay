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

// openssl prime -generate -bits 1024

func main() {
	var plainText = "ACLOUD"

	// 1. Choose two distinct prime numbers, p and q
	// 4. Choose any number 1 < e < 780 that is coprime to 780.
	var p, q, e big.Int
	p.SetString("149167516042398723525996758773222137472281596223800479849551540351321034501792773631372499057699531070744472669665053159335520543433504855361286774214076723836802002410302647566892453660983455011781051513658262625786070114418438980094403935311556559954980492420486666809425360953717676668351593906539161639037", 10)
	q.SetString("144085140269927785077965383593233311382753811804186408035700201265790515428737655822413298633470907606520085298502109809391457262279140832852147401148057790649935817654787884130381388430766729931091479265582830015230104977199529156176249346200528785749827308279438469063610414582644920648362498649191416059531", 10)
	e.SetString("164980429110169948421905422637425619464250398994123121994248535750968988933737559448128821647637771319856056232350463890336363482800239649343244642618554310520539501215821564433169870438486644707957411182704673414335171493308674958494165428598801508076047085838657336277080545578857427910086984801408944113217", 10)

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
