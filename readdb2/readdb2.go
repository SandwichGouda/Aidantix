package main

import (
	"log"
	"os"
	"strings"
	"fmt"
)

type Word struct {
	Class     string
	Frequency int
	Label     string
}

func ImportDatabase() []string {

	// db, err := os.ReadFile("readdb2/database/scraped_words.csv") // returns a []byte
	db, err := os.ReadFile("database/scraped_words.csv") // returns a []byte

	if err != nil {
		log.Fatal("readdb2.go : Failed to read database \n", err)
	}
	
	return strings.Split(string(db),"\n")
}

func main() {
	fmt.Println(ImportDatabase()[:10])
}