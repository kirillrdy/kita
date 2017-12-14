package error

import (
	"log"
)

//TODO move to util
func Crash(err error) {
	if err != nil {
		log.Panic(err)
	}
}
