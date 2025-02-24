package scrape

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func ScrapeDay() int {
	resp, err := http.Get("https://cemantix.certitudes.org/")
	if err != nil {
		fmt.Println("scrape.go : ", err)
		return 0
	}

	var (
		method1_criticalfail bool
		method2_criticalfail bool
	)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("scrape.go : ", err)
		return 0
	}
	page := string(body)

	i_1 := strings.Index(page, "data-puzzle-number=")
	if i_1 == -1 {
		fmt.Println(`"scrape.go : data-puzzle-number=" not found on page"`)
	} else {
		i_1 += 20
	}

	day_1, err := strconv.Atoi(page[i_1 : i_1+4])
	if err != nil {
		fmt.Println("scrape.go : Failed to convert string page[i_1 : i_1+4] = ", page[i_1:i_1+4], "to int.")
		method1_criticalfail = true
	}

	i_2 := strings.Index(page, `id="puzzle-num">`)
	if i_2 == -1 {
		fmt.Println(`scrape.go : id="puzzle-num"> not found on page"`)
		method2_criticalfail = true
	} else {
		i_2 += 16
	}

	day_2, err := strconv.Atoi(page[i_2 : i_2+4])
	if err != nil {
		fmt.Println("scrape.go : Failed to convert string page[i_2 : i_2+4] = ", page[i_2:i_2+4], "to int.")
		method2_criticalfail = true
	}

	switch {
	case method1_criticalfail && method2_criticalfail:
		fmt.Println("scrape.go : both methods failed")
		return 0
	case method1_criticalfail:
		fmt.Println("scrape.go : method 1 failed, returning day_2")
		return day_2
	case method2_criticalfail:
		fmt.Println("scrape.go : method 2 failed, returning day_1")
		return day_1
	default:
		if day_1 != day_2 {
			fmt.Println("scrape.go : day_1 =", day_1, " != ", day_2, " = day_2 ; returning day_1")
		}
		return day_1
	}
}
