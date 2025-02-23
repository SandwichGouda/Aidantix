package scrape

import (
	// "fmt"
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

	i_1 := strings.Index(page, "data-puzzle-number=")
	if i_1 == -1 {
		log.Fatal(`"data-puzzle-number=" not found on page"`)
	}
	i_1 += 20

	day_1, err := strconv.Atoi(page[i_1 : i_1+4])
	if err != nil {
		log.Fatal("Failed to convert string page[i_1 : i_1+4] = ", page[i_1:i_1+4], "to int.")
	}

	i_2 := strings.Index(page, `id="puzzle-num">`)
	if i_2 == -1 {
		log.Fatal(`id="puzzle-num"> not found on page"`)
	}
	i_2 += 16

	day_2, err := strconv.Atoi(page[i_2 : i_2+4])
	if err != nil {
		log.Fatal("Failed to convert string page[i_2 : i_2+4] = ", page[i_2:i_2+4], "to int.")
	}

	if day_1 != day_2 {
		log.Fatal("day_1 =", day_1, " != ", day_2, " = day_2")
	}
	return day_1
}
