package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Korek struct {
	hits       int
	lastPlayer string
}

func main() {
	lanjut := make(chan *Korek)
	kalah := make(chan *Korek)

	breakpoint := 10
	rand.Seed(time.Now().UTC().UnixNano())

	for {
		players := []string{"Player 1", "Player 2", "Player 3"}
		for _, p := range players {
			go ambilKorek(p, breakpoint, lanjut, kalah)
		}

		lanjut <- new(Korek)

		select {
		case d := <-kalah:
			fmt.Println(d.lastPlayer, "kalah pada hit ke", d.hits)
			return
		}
	}
}

func ambilKorek(player string, breakpoint int, lanjut, kalah chan *Korek) { //(isEnd bool) {

	min := 1
	max := 100
	for {
		select {
		case k := <-lanjut:

			nilaiKorek := rand.Intn(max-min) + min
			time.Sleep(500 * time.Millisecond)

			k.hits++
			k.lastPlayer = player
			fmt.Println("korek ada di", k.lastPlayer, "pada hit ke", k.hits, "dan mempunyai nilai", nilaiKorek)
			result := nilaiKorek % breakpoint

			if result == 0 {
				kalah <- k
				return
			}
			lanjut <- k
		}
	}
}
