package validators

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/google/uuid"
)

func IsValidPhoneNumber(phone_number string) bool {
	e164Regex := `^\+[1-9]\d{1,14}$`
	re := regexp.MustCompile(e164Regex)
	phone_number = strings.ReplaceAll(phone_number, " ", "")

	return re.Find([]byte(phone_number)) != nil
}

func IsValidPass(passport string) bool {
	e164Regex := `^[1-9]{10}$`
	re := regexp.MustCompile(e164Regex)
	passport = strings.ReplaceAll(passport, " ", "")

	return re.Find([]byte(passport)) != nil
}

func IsValidLen(name string, l int) bool {
	return len(name) > l
}

func IsValidLenEq(val string, l int) bool {
	return utf8.RuneCountInString(val) == l
}

func IsValidGuid(guid uuid.UUID) bool {
	err := uuid.Validate(guid.String())
	return err == nil && guid != uuid.Nil
}

func IsValidSum(sum int) bool {
	return sum > 0
}
