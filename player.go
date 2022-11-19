package main

import (
	"fmt"
	"strings"
)

func playerTurn(allPlayers []*Hand, deck Deck) ([]*Hand, Deck) {
	var move string
	var card Card

	for i := 0; i < len(allPlayers)-1; i++ {
		iter := 0
		total, _ := getTotal(*allPlayers[i])
		fmt.Printf("\nPlayer %v's turn (Total: %v)\n", i+1, total)
		for {
			if iter > 0 {
				fmt.Printf("Player %v: %v\n", i+1, printTotal(*allPlayers[i]))
			}
			if allPlayers[i].checkBlackjack() {
				fmt.Println("Blackjack!")
				break
			}
			if allPlayers[i].checkBust() {
				fmt.Println("Bust.")
				fmt.Println()
				allPlayers[i].bust = true
				break
			}
			fmt.Println("Hit or stand?")
			iter++
			fmt.Scanf("%s", &move)
			move = strings.ToLower(move)
			if move == "hit" {
				card, deck = drawCard(deck)
				printHit(card, "Player")
				allPlayers[i].hand = append(allPlayers[i].hand, card)
				continue
			}
			if move == "stand" {
				fmt.Println()
				break
			}
			fmt.Println("Invalid entry")
		}
	}
	return allPlayers, deck
}
