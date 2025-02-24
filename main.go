package main

import (
	"aidantix/readdb"
	"aidantix/req"
	"aidantix/scrape"
	"aidantix/scrape_goquery"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Response struct {
	Word    string
	Rank    int
	Score   float64
	Unknown bool
}

func real_main() {

	var day int

	day = scrape.ScrapeDay()
	if day == 0 {
		fmt.Println("scrape.ScrapeDay() returned 0. Moving onto scrape_goquery.ScrapeDay()")
	}

	day = scrape_goquery.ScrapeDay()
	if day == 0 {
		fmt.Print("scrape_goquery.ScrapeDay() returned 0. Please type in the puzzle day number : ")
		reader := bufio.NewReader(os.Stdin)
		nb, _ := reader.ReadString('\n')
		nb = nb[:len(nb)-1]

		var err error
		day, err = strconv.Atoi(nb)

		for err != nil {
			fmt.Print("Input convertion to int type failed. Please type in the puzzle day number : ")
			nb, _ = reader.ReadString('\n')
			nb = nb[:len(nb)-1]
			day, err = strconv.Atoi(nb)
		}
	}

	fmt.Println("main : day", day)

	fmt.Println(req.Request("valeur", day))

	w := db1.ImportDatabase()

	for i, wd := range w {
		fmt.Println(i, wd.Label)
	}

	fmt.Println(w)

	// db1.Write()
}

func main() {

	w := db1.ImportDatabase()

	for i, wd := range w {
		fmt.Println(i, wd.Label)
	}

	fmt.Println(w)
}
