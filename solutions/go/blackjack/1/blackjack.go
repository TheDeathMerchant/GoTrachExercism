package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    switch card {
        case "ace":
        	return 11
        case "two":
        	return 2
        case "three":
        	return 3
        case "four":
        	return 4
        case "five":
        	return 5
        case "six":
        	return 6
        case "seven":
        	return 7
        case "eight":
        	return 8
        case "nine":
        	return 9
        case "ten", "jack", "queen", "king":
        	return 10
        default:
        	return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	s1 := ParseCard(card1)
    s2 := ParseCard(card2)
    s3 := ParseCard(dealerCard)

    if s1 == 11 && s2 == 11 {
        return "P"
    } else if s1 + s2 == 21 && s3 != 11 && s3!= 10 {
        return "W"
    } else if s1 + s2 >=17 &&  s1 + s2 <=20 {
        return "S"
    } else if  s1 + s2 >=12 &&  s1 + s2 <=16 && s3 < 7 {
        return "S"
    } else if s1 + s2 >=12 &&  s1 + s2 <=16 && s3 >=7 {
        return "H"
    } else if s1 + s2 <= 11 {
        return "H"
    }

    return "S"
}
