package main

import "log"

func CheckError(err error, args ...string) {
	if err != nil {
		log.Println("Error:", err)
		log.Println("Stack trace:")
		log.Printf("%+v\n", err)
		if len(args) > 0 {
			log.Printf("Aditional error info: %v", args)
		}
	}
}
