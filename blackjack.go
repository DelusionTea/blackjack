package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	//panic("Please implement the ParseCard function")

	if card == "jack" || card == "queen" || card == "king" || card == "ten" {
		return 10
	}
	switch card {
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	case "ace":
		return 11
	default:
		return 0
	}

}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	//panic("Please implement the IsBlackjack function")
	if (ParseCard(card1) + ParseCard(card2)) == 21 {
		return true
	}
	return false
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	//panic("Please implement the LargeHand function")
	//	As the `LargeHand` function is only called for hands with a value larger than 20,
	//	there are only 2 different possible hands: A **BlackJack** with a total value of `21` and **2 Aces** with a total value of `22`.
	//	- The function should check [if][if_statement] `isBlackJack` is `true` and return "P" otherwise.
	//	- If `isBlackJack` is `true`, the dealerScore needs to be checked for being lower than 10.
	//	[If][if_statement] it is lower, return "W" otherwise "S"
	if isBlackjack {
		if dealerScore < 10 {
			return "W"
		} else {
			return "S"
		}
	} else {
		return "P"
	}
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	//panic("Please implement the SmallHand function")
	if handScore >= 17 || (handScore >= 12 && dealerScore <= 6) {
		return "S"
	}
	return "H"
}
