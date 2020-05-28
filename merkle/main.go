package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"time"
)

type block struct {
	ID           uint
	PreviousHash string
	Timestamp    time.Time
	TxRoot       string
	Nonce        string
}

type node struct {
	Hash string
}

func shahash(input string) string {
	h := sha512.New()
	h.Write([]byte(input))
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}

func hash(input string) string {
	return shahash(shahash(input))
}

func main() {
	genesis := hash("Levi 2020-05-11")
	fmt.Println(genesis)
}
