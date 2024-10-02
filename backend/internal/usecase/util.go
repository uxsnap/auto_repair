package usecase

import (
	"log"
	"regexp"
	"strings"
)

func ServiceLogger(prefix string, cb func() interface{}) interface{} {
	log.Printf("calling %v usecase", prefix)
	res := cb()

	return res
}

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
