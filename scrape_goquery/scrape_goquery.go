package scrape_goquery

import (
	// "fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func ScrapeDay() int {
	resp, err := http.Get("https://cemantix.certitudes.org/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	page := string(body)

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(page))
	if err != nil {
		log.Fatal(err)
	}

	var day_1, day_2 *int

	// Find all <script> tags with id="main_script"
	doc.Find("script[data-puzzle-number]").Each(func(index int, element *goquery.Selection) {
		// Get the value of the data-puzzle-number attribute
		data_puzzle_number, exists := element.Attr("data-puzzle-number")
		if exists {
			d, err := strconv.Atoi(data_puzzle_number)
			if err != nil {
				log.Fatal("Unable to convert data-puzzle-number =", data_puzzle_number, "to int")
			}
			day_1 = &d
		}
	})

	doc.Find("b#puzzle-num").Each(func(index int, element *goquery.Selection) {
		tagContent := element.Text()
		d, err := strconv.Atoi(tagContent)
		if err != nil {
			log.Fatal(`Unable to convert "puzzle-num" tag content `, tagContent, " to int")
		}
		day_2 = &d
	})

	if day_1 == nil {
		log.Fatal("day_1 is nil pointer in scrape_go : retrieveing day_1 failed")
	}

	if day_2 == nil {
		log.Fatal("day_2 is nil pointer in scrape_go : retrieveing day_2 failed")
	}

	// fmt.Println(*day_1, *day_2)

	return *day_1
}

/*
TEST & DEBUG PAGE

page := `
<!DOCTYPE html>
<html>
<head>
	<script id="main_script">console.log("Script 1");</script>
	<script>console.log("Script 2");</script>
	<script id="script" type="module" src="https://static.certitudes.org/html/cemantix.js" data-puzzle-number="1089" data-utc-time="1740265200" data-app="cemantix"></script>
	h3>Jour nº<b id="puzzle-num">1089</b></h3>
</head>
<body>
</body>
</html>`
*/
