package assignment

import (
	"fmt"
	"math/rand"
	"time"
)

func ChannelKorekApiBernyanyi() {
	var (
		breakPoint   = 11
		no           = 1
		player       = []string{"A", "B", "C", "D"}
		chanPlayer   = make(chan string)
		playerName   string
		randonNumber int
	)

	rand.Seed(time.Now().UnixNano())
	for {
		for _, p := range player {
			go func(name string) {
				chanPlayer <- name
			}(p)
		}

		playerName = <-chanPlayer
		randonNumber = rand.Intn(100)
		fmt.Println("korek ada di player", playerName, "pada hit ke", no, "dan mempunyai nilai", randonNumber)
		if randonNumber%breakPoint == 0 {
			fmt.Println("player", playerName, "kalah pada hit ke", no)
			break
		}

		no++
	}
}
