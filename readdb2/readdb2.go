package readdb2

import (
	"log"
	"os"
	"strings"
)

func ImportDatabase() []string {

	db, err := os.ReadFile("readdb2/database/words.csv") // returns a []byte

	if err != nil {
		log.Fatal("readdb2.go : Failed to read database \n", err)
	}

	return strings.Split(string(db),"\n")
}