package main

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

type email struct {
	ID        int
	Hash      string
	SpamScore int
}

var s rand.Source
var r *rand.Rand

func init() {
	s = rand.NewSource(time.Now().UnixNano())
	r = rand.New(s)
}

func emailProcessor(emails []email) {
	em := make(chan email)
	hashed := make(chan email)
	go emailHasher(emails, em)
	go emailSpammer(em, hashed)
	emailSaver(hashed)
}

func emailHasher(emails []email, ch chan email) {
	for _, e := range emails {
		rnd := r.Int() % 10
		rnds := time.Duration(rnd)
		time.Sleep(time.Second * rnds)

		color.Blue("Hasher\tEmail\t%d\tSlept\t%d\n", e.ID, rnds)

		ch <- e
	}
	close(ch)
}

func emailSpammer(chin chan email, chout chan email) {
	for e := range chin {
		rnd := r.Int() % 10
		rnds := time.Duration(rnd)
		time.Sleep(time.Second * rnds)
		color.Red("Spammer\tEmail\t%d\tSlept\t%d\n", e.ID, rnds)

		chout <- e
	}
	close(chout)
}

func emailSaver(chin chan email) {
	for e := range chin {
		rnd := r.Int() % 10
		rnds := time.Duration(rnd)
		time.Sleep(time.Second * rnds)
		color.Green("Saver\tEmail\t%d\tSlept\t%d\n", e.ID, rnds)
	}
}

// Do not communicate by sharing memory; instead, share memory by communicating.
// https://blog.golang.org/codelab-share
func main() {
	var emails = []email{}

	for i := 0; i < 10; i++ {
		emails = append(emails, email{ID: i})
	}

	emailProcessor(emails)
}
