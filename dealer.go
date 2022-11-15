package main

import "fmt"

func dealerTurn(dealer *Hand, deck Deck) *Hand {
	//check score, on 16 or less hit on soft 17 hit
	var card Card
	for {
		total, soft := getTotal(*dealer)
		if (total == 17 && !soft) || total > 16 {
			break
		}
		card, deck = drawCard(deck)
		dealer.hand = append(dealer.hand, card)
		fmt.Println(printTotal(*dealer))
	}
	return dealer
}