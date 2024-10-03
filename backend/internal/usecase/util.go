package usecase

import (
	"log"
)

func ServiceLogger(prefix string, cb func() interface{}) interface{} {
	log.Printf("calling %v usecase", prefix)
	res := cb()

	return res
}
