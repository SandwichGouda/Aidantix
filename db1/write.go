package db1

import (
	"fmt"
	"log"
	"os"
)

type Response struct {
	Word    string
	Rank    int
	Score   float64
	Unknown bool
}

func Write(resp Response) {

	// err := os.WriteFile("db1/result", []byte("hello\ngo\n"), 0666)

	f, err := os.OpenFile("db1/result", os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	if _, err := f.Write([]byte("appended some data\n")); err != nil {
		fmt.Println("Here we are")
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	// d2 := []byte{115, 111, 109, 101, 10}
	// n2, err := f.Write(d2)
	// check(err)
	// fmt.Printf("wrote %d bytes\n", n2)

	// n3, err := f.WriteString("writes\n")
	// check(err)
	// fmt.Printf("wrote %d bytes\n", n3)

	// f.Sync()

}
