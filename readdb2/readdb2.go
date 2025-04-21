package readdb2

import (
	"log"
	"os"
	"strings"
	// "fmt"
)

func ImportDatabase() []string {

	db, err := os.ReadFile("readdb2/database/words.csv") // returns a []byte

	if err != nil {
		log.Fatal("readdb2.go : Failed to read database \n", err)
	}
 
	str_db_with_freq := strings.Split(string(db),"\n")
	
	str_db := make([]string,10000)

	for i,wd := range str_db_with_freq {
		str_db[i] = strings.Split(wd,",")[1]
	}

	return str_db
}