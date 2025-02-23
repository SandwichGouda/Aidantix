package req

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Response struct {
	Word    string
	Rank    int
	Score   float64
	Unknown bool
}

func Request(word string, day int) *Response {

	day_str := strconv.Itoa(day)

	requestURL := "https://cemantix.certitudes.org/score?n=" + day_str
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

	// fmt.Printf("Requested : %s | Response body: %s", word, resBody)

	// responses are either of the form :
	// {"p":214,"s":0.1799,"v":13135} if word is 214th
	// {"s":-0.0078,"v":13135} if word is >1000th
	// {"e":"Je ne connais pas le mot <i>unknownword</i>","v":13135} if word is unknown

	var preresp map[string]interface{}

	resp := &Response{
		Word: word,
	}

	if err := json.Unmarshal(resBody, &preresp); err != nil {
		log.Fatal(err)
	}

	_, ok_r := preresp["r"]

	if ok_r { // The key "r" is in the server's response : The day is wrong
		log.Fatal("The key 'r') is in the server's response. The request day is likely wrong")
	}

	elem_e, ok_e := preresp["e"]

	if ok_e { // The key "e" is in the server's response : It doesn't know the word
		// fmt.Println("Word", word, "is unknown :", elem_e)
		resp = &Response{
			Word:    word,
			Rank:    0,
			Score:   0.0,
			Unknown: true,
		}
		return resp
	}

	// If we got here, the word should've been recognised.
	score, ok_s := preresp["s"].(float64)

	if !ok_s {
		log.Fatal("Score conversion failed. Server's response :", string(resBody))
	}

	_, ok_p := preresp["p"]

	if ok_p { // The key "p" is in the server's response : We hit 1000
		// fmt.Println("Word", word, "hit 1000 ! (", elem_p, ")")

		rank, ok_p := preresp["p"].(float64)

		if !ok_p {
			log.Fatal("Rank conversion failed. Server's response :", string(resBody))
		}

		resp = &Response{
			Word:    word,
			Rank:    int(rank),
			Score:   score,
			Unknown: false,
		}
		return resp
	}

	// Word did not hit 1000, or server responded something else
	// fmt.Println("Word", word, " did not hit 1000")
	resp = &Response{
		Word:    word,
		Rank:    0,
		Score:   score,
		Unknown: false,
	}
	return resp
}
