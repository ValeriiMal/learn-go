package blackjack

import (
	"slices"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "eight":
		return 8
	case "two":
		return 2
	case "nine":
		return 9
	case "three":
		return 3
	case "ten":
		return 10
	case "four":
		return 4
	case "jack":
		return 10
	case "five":
		return 5
	case "queen":
		return 10
	case "six":
		return 6
	case "king":
		return 10
	case "seven":
		return 7
	case "other":
		return 0
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	dealerCardStand := []string{"ace", "jack", "queen", "kind"}
	cardSum := ParseCard(card1) + ParseCard(card2)
	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case cardSum == 21 && !slices.Contains(dealerCardStand, dealerCard):
		return "W"
	case cardSum == 21 && slices.Contains(dealerCardStand, dealerCard):
		return "S"
	case cardSum >= 17 && cardSum <= 20:
		return "S"
	case cardSum >= 12 && cardSum <= 16 && ParseCard(dealerCard) < 7:
		return "S"
	case cardSum >= 12 && cardSum <= 16 && ParseCard(dealerCard) >= 7:
		return "H"
	case cardSum <= 11:
		return "H"
	default:
		return "P"
	}
}
