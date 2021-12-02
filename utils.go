package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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

func fetch(url string) []string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: readFile(".cookie")})

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Fields(string(body))

	return input
}

func strArrToIntArr(arr []string) []int {
	var result []int
	for i := 0; i < len(arr); i++ {
		element, err := strconv.Atoi(arr[i])
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, element)
	}
	return result
}

func sum(input []int) int {
	sum := 0

	for _, i := range input {
		sum += i
	}

	return sum
}

func strToInt(str string) int {
	num, err := strconv.Atoi(str)

	if err != nil {
		log.Fatal(err)
	}
	return num
}
