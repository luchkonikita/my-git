package errors

import (
	"log"
)

// CheckError - handle error if exists.
func CheckError(err interface{}) {
	if err != nil {
		log.Fatal(err)
	}
}
