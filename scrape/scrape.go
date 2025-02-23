package scrape

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://cemantix.certitudes.org/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
