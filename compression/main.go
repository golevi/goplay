//
// compression
//
// references used to write this
// - https://en.wikipedia.org/wiki/Huffman_coding
// - https://golang.org/ref/spec
// - https://golang.org/pkg/bytes/
// - https://www.tutorialspoint.com/go/go_bitwise_operators.htm
// - https://web.stanford.edu/class/archive/cs/cs106b/cs106b.1126/handouts/220%20Huffman%20Encoding.pdf
//
package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

// const input string = "hello world hello world hello go"
const input string = `
Test compression is a technique used to reduce the time and cost of testing integrated circuits. The first ICs were tested with test vectors created by hand. It proved very difficult to get good coverage of potential faults, so Design for testability (DFT) based on scan and automatic test pattern generation (ATPG) were developed to explicitly test each gate and path in a design. These techniques were very successful at creating high-quality vectors for manufacturing test, with excellent test coverage. However, as chips got bigger the ratio of logic to be tested per pin increased dramatically, and the volume of scan test data started causing a significant increase in test time, and required tester memory. This raised the cost of testing.

Test compression was developed to help address this problem. When an ATPG tool generates a test for a fault, or a set of faults, only a small percentage of scan cells need to take specific values. The rest of the scan chain is don't care, and are usually filled with random values. Loading and unloading these vectors is not a very efficient use of tester time. Test compression takes advantage of the small number of significant values to reduce test data and test time. In general, the idea is to modify the design to increase the number of internal scan chains, each of shorter length. These chains are then driven by an on-chip decompressor, usually designed to allow continuous flow decompression where the internal scan chains are loaded as the data is delivered to the decompressor. Many different decompression methods can be used.[1] One common choice is a linear finite state machine, where the compressed stimuli are computed by solving linear equations corresponding to internal scan cells with specified positions in partially specified test patterns. Experimental results show that for industrial circuits with test vectors and responses with very low fill rates, ranging from 3% to 0.2%, the test compression based on this method often results in compression ratios of 30 to 500 times.[2]

With a large number of test chains, not all the outputs can be sent to the output pins. Therefore, a test response compactor is also required, which must be inserted between the internal scan chain outputs and the tester scan channel outputs. The compactor must be synchronized with the data decompressor, and must be capable of handling unknown (X) states. (Even if the input is fully specified by the decompressor, these can result from false and multi-cycle paths, for example.) Another design criteria for the test result compressor is that it should give good diagnostic capabilities, not just a yes/no answer.
`

func main() {

	var symbols = make(map[byte]int)

	for _, i := range input {
		_, ok := symbols[byte(i)]
		if ok {
			symbols[byte(i)]++
		} else {
			symbols[byte(i)] = 1
		}
	}
	var binaryMap = make(map[byte]int)
	var k int = 0
	for i := range symbols {
		binaryMap[byte(i)] = k
		k++
	}

	// fmt.Println(symbols)
	// fmt.Println(binaryMap)

	var a = 0
	var b byte
	var bs []byte

	for x, i := range input {
		if x%2 == 0 {
			// Set the right half of the byte 0000 xxxx to the byte value from
			// the binaryMap converting it to a byte.
			a = binaryMap[byte(i)]
		} else {
			// Move the data in the right half of the byte to the left half to
			// make room for the next value
			// before 0000 xxxx
			// after  xxxx 0000
			a = a << 4
			// Add the new right part of the byte
			// xxxx xxxx
			a = a | binaryMap[byte(i)]

			// Convert it to a byte
			b = byte(a)

			// Append it to the byte array
			bs = append(bs, b)
		}
	}
	fmt.Println(len(bs) * 8)
	fmt.Println(len(input) * 8)

	ratio := ((float32(len(bs) * 8)) / float32((len(input) * 8)) * 100)
	fmt.Printf("%f%%\n", ratio)

	fmt.Println(hex.EncodeToString(bs))
	fmt.Println(base64.StdEncoding.EncodeToString(bs))

	f, _ := os.Create("tmp")
	defer f.Close()
	f.Write(bs)
}
