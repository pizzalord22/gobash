package main

import (
	"github.com/pkg/errors"
	"log"
)

// takes at least 1 error as argument
// it can also take 0 or more strings as additional info
func CheckError(err error, args ...string) {
	if err != nil {
		log.Println("Error:", err.Error())
		log.Println("stack trace,\n", errors.Errorf("%+v", err))
		if len(args) > 0 {
			log.Printf("Aditional error info: %v\n", args)
		}
	}
}
