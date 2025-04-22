package readdb2

import (
	"log"
	"os"
	"strings"
	"fmt"
)

func ImportDatabase() []string {

	db, err := os.ReadFile("readdb2/database/words.csv") // returns a []byte

	if err != nil {
		log.Fatal("readdb2.go : Failed to read database \n", err)
	}

	strdb := strings.Split(string(db),"\n")

	for i,wd := range strdb {
		strdb[i] = wd[:len(wd)-1]
	}

	fmt.Println("Words :",strdb[0],strdb[1],strdb[2])

	fmt.Println("Bytes :",[]byte(strdb[0]),[]byte(strdb[1]),[]byte(strdb[2]))

	return strdb
}