package main

import (
	"io/ioutil"
	"log"
)

func readFile(filename string) string {
	content, err := ioutil.ReadFile(".cookie")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string
	text := string(content)
	return text
}
