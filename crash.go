package kita

import (
	"log"
)

func Crash(err error) {
	if err != nil {
		log.Panic(err)
	}
}
