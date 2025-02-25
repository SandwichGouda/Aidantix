package readdb1

import (
	"encoding/json"
	"log"
	"os"
)

type Word struct {
	Class     string
	Frequency int
	Label     string
}

func ImportDatabase() []string {

	db, err := os.ReadFile("readdb1/database/words.json") // returns a []byte

	if err != nil {
		log.Fatal("readdb1.go : Failed to read database \n", err)
	}

	// testdb := `[{"Class":"dét.","Frequency":1050561,"Label":"le"},{"Class":"prép.","Frequency":862100,"Label":"de"}]`
	var words []Word
	if err := json.Unmarshal(db, &words); err != nil {
		log.Fatal("readdb1.go : Failed to unmarshal database \n", err)
	}

	words_str := make([]string, len(words), len(words))

	for i,w := range words {
		words_str[i] = w.Label
	}

	return words_str
}
