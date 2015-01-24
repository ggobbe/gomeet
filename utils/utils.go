package utils

import "log"

func CheckErrorMsg(err error, message string) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
