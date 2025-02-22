package req

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Request(word string) {

	requestURL := "https://cemantix.certitudes.org/score?n=1088"
	body := strings.NewReader("word=" + word)

	req, err := http.NewRequest("POST", requestURL, body)

	if err != nil {
		fmt.Printf("Could not create request: %s\n", err)
		os.Exit(1)
	}

	req.Header.Set("origin", "https://cemantix.certitudes.org")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("Error making http request: %s\n", err)
		os.Exit(1)
	}

	// fmt.Println("client: got response!")                 // To be commented when finished
	// fmt.Println("client: status code: ", res.StatusCode) // To be commented when finished

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Requested : %s | Response body: %s", word, resBody)
}
