package scrape_goquery

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func ScrapeDay() int {
	resp, err := http.Get("https://cemantix.certitudes.org/")
	if err != nil {
		fmt.Println("scrape_goquery.go : err")
		return 0
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("scrape_goquery.go : ", err)
		return 0
	}
	page := string(body)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(page))
	if err != nil {
		fmt.Println("scrape_goquery.go : ", err)
		return 0
	}

	var day_1, day_2 *int

	doc.Find("script[data-puzzle-number]").Each(func(index int, element *goquery.Selection) {
		data_puzzle_number, exists := element.Attr("data-puzzle-number")
		if exists {
			d, err := strconv.Atoi(data_puzzle_number)
			if err != nil {
				fmt.Println("scrape_goquery.go : Unable to convert data-puzzle-number =", data_puzzle_number, "to int")
			}
			day_1 = &d
		}
	})

	doc.Find("b#puzzle-num").Each(func(index int, element *goquery.Selection) {
		tagContent := element.Text()
		d, err := strconv.Atoi(tagContent)
		if err != nil {
			fmt.Println(`scrape_goquery.go : Unable to convert "puzzle-num" tag content `, tagContent, " to int")
		}
		day_2 = &d
	})

	switch {
	case day_1 == nil && day_2 == nil:
		fmt.Println("scrape_goquery.go : day_1 and day_2 are both nil pointers in scrape_go : retrieveing day_1 & day_2 failed")
		return 0
	case day_2 == nil:
		fmt.Println("scrape_goquery.go : day_2 is nil pointer in scrape_go : retrieveing day_2 failed ; but apparently day_1 did not")
		return *day_1
	case day_1 == nil:
		fmt.Println("scrape_goquery.go : day_1 is nil pointer in scrape_go : retrieveing day_1 failed ; but apparently day_2 did not")
		return *day_2
	default:
		return *day_1
	}
}