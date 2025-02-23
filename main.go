package main

import (
	"aidantix/req"
	"aidantix/scrape"
	"aidantix/scrape_goquery"
	"fmt"
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
		fmt.Println("scrape_goquery.ScrapeDay() returned 0. Please type in the puzzle day number :")

	}

	fmt.Println("main : day", day)

	fmt.Println(req.Request("valeur", day))

	w := db1.ImportDatabase()

	for i, wd := range w {
		fmt.Println(i, wd.Label)
	}

	fmt.Println(w)

	db1.Write()
}

func main() {

	day := scrape_goquery.ScrapeDay()

	fmt.Println("main : day", day)
}
