package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

func main() {

	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<script id="main_script">console.log("Script 1");</script>
		<script>console.log("Script 2");</script>
		<script id="script" type="module" src="https://static.certitudes.org/html/cemantix.js" data-puzzle-number="1089" data-utc-time="1740265200" data-app="cemantix"></script>
		<h3>Jour nº<b id="puzzle-num">1089</b></h3>
	</head>
	<body>
	</body>
	</html>`

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	// vat day_1, day_2 *int

	// Find all <script> tags with id="main_script"
	doc.Find("b#puzzle-num").Each(func(index int, element *goquery.Selection) {
		tagContent := element.Text()
		fmt.Printf("Script %d: %s\n", index+1, tagContent)
	})
}
