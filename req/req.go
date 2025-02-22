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
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	req.Header.Set("origin", "https://cemantix.certitudes.org")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("client: got response!")
	fmt.Println("client: status code: ", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", resBody)
}
