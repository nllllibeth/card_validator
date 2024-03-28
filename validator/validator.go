package validator

import (
	"strings"
)

func IsValidCreditCardNumber(cardNumber string) bool {
	cardNumber = strings.ReplaceAll(cardNumber, " ", "")
	cardNumber = strings.ReplaceAll(cardNumber, "-", "")

	for _, char := range cardNumber {
		if char < '0' || char > '9' {
			return false
		}
	}
	if len(cardNumber) < 13 || len(cardNumber) > 19 {
		return false
	}

	// the Luhn's algorithm 
	oddSum := 0
	evenSum := 0
	for i := 0; i < len(cardNumber); i++ {
		if i%2 != 0 {
			oddSum += int(cardNumber[i] - '0')
		} else {
			digit := int(cardNumber[i]-'0') * 2
			if digit > 9 {
				digit = (digit % 10) + (digit / 10)
			}
			evenSum += digit
		}
	}
	sum := oddSum + evenSum

	return sum %10 == 0
}