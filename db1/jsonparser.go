package db1

import (
	"encoding/json"
	"log"
)

type Word struct {
	Class     string
	Frequency int
	Label     string
}

func ImportDatabase() []Word {

	db := `[{"Class":"dét.","Frequency":1050561,"Label":"le"},{"Class":"prép.","Frequency":862100,"Label":"de"}]`
	var words []Word
	if err := json.Unmarshal([]byte(db), &words); err != nil {
		log.Fatal(err)
	}

	return words
}
