package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
)

type (
	// Struct for a array of quotes
	Quotes struct {
		Quotes []Quote `json:"quotes"`
	}

	// Struct for a single quote
	Quote struct {
		Author string `json:"author"`
		Quote  string `json:"quote"`
	}
)

func main() {
	// read the external json file and returns the JSON-encoded data as "f"
	f, err := os.ReadFile("./quotes.json")
	if err != nil {
		log.Fatalf("Unable to read file due to %s\n", err)
	}

	var quotes Quotes

	// we use json.Unmarshal to deserialize "f" and store the result it in "quotes" variable
	err = json.Unmarshal(f, &quotes.Quotes)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}

	quotesCount := len(quotes.Quotes)

	// random number between 0 and (quotesCount-1)
	randomIndex := rand.Intn(quotesCount-1) + 0

	fmt.Printf("%s - %s \n", quotes.Quotes[randomIndex].Quote, quotes.Quotes[randomIndex].Author)
}
