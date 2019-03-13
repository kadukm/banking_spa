package utils

import (
	"regexp"
	"strings"
	"unicode"
)

var emailPattern *regexp.Regexp = regexp.MustCompile(`\w+@\w+\.\w+`)
var expiresPattern *regexp.Regexp = regexp.MustCompile(`(0[1-9]|1[0-2])/(19|[23]\d|2[0-5])`)

func IDIsRight(id string) bool {
	return id == ""
}

func INNIsRight(inn string) bool {
	innLength := len(inn)
	return (innLength == 10 || innLength == 12) && stringContainsOnlyDigits(inn)
}

func BIKIsRight(bik string) bool {
	bikLength := len(bik)
	return bikLength == 9 && stringContainsOnlyDigits(bik)
}

func AccountNumberIsRight(accountNumber string) bool {
	accountNumberLength := len(accountNumber)
	return accountNumberLength == 20 && stringContainsOnlyDigits(accountNumber)
}

func ForWhatIsRight(forWhat string) bool {
	return strings.Contains(forWhat, "без НДС") ||
		strings.Contains(forWhat, "НДС 10%") ||
		strings.Contains(forWhat, "НДС 18%")
}

func ValueIsRight(value int) bool {
	return value >= 1000 && value <= 75000
}

func PhoneIsRight(phone string) bool {
	firstChar := phone[0]
	return firstChar == '+' && stringContainsOnlyDigits(phone[1:])
}

func EmailIsRight(email string) bool {
	return emailPattern.MatchString(email)
}

func CardNumberIsRight(accountNumber string) bool {
	accountNumberLength := len(accountNumber)
	return accountNumberLength == 16 && stringContainsOnlyDigits(accountNumber)
}

func CardExpiresIsRight(expires string) bool {
	return expiresPattern.MatchString(expires)
}

func CardCvcIsRight(cvc int) bool {
	return cvc >= 100 && cvc <= 999
}

func CommentIsRight(comment string) bool {
	commentLength := len(comment)
	return commentLength <= 150
}

func stringContainsOnlyDigits(value string) bool {
	for _, c := range value {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
