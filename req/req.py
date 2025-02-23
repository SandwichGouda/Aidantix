package req

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Response struct {
	Word    string
	Rank    int
	Score   float64
	Unknown bool
}

type PreResponseWithRank struct {
	p int
	s float64
	v int
}

type PreResponseWithoutRank struct {
	s float64
	v int
}

type PreResponseError struct {
	e string
	v int
}

func Request(word string) *Response {

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

	// responses are either of the form :
	// {"p":214,"s":0.1799,"v":13135} if word is 214th
	// {"s":-0.0078,"v":13135} if word is >1000th
	// {"e":"Je ne connais pas le mot <i>unknownword</i>","v":13135} if word is unknown

	preresprank := &PreResponseWithRank{
		p: 0,
		s: 5,
		v: 0,
	}

	prerespnorank := &PreResponseWithoutRank{
		s: 0,
		v: 0,
	}

	preresperror := &PreResponseError{
		e: "",
		v: 0,
	}

	// resp.Word = word

	if err := json.Unmarshal(resBody, preresprank); err != nil {
		// Word is not "recognised but outside first 1000"

		if err = json.Unmarshal(resBody, prerespnorank); err != nil {
			// Word is not "recognised and in first 1000"

			if err = json.Unmarshal(resBody, preresperror); err != nil {
				// Word is not "unrecognised"

				log.Fatal(err)
			} else {
				fmt.Println(word, "is an unknown word for cemantix.")

				return &Response{
					Word:    word,
					Rank:    0,
					Score:   0.0,
					Unknown: true,
				}
			}

		} else {

			fmt.Println(word, "did not hit 1000 :", string(resBody))
			return &Response{
				Word:    word,
				Rank:    0,
				Score:   prerespnorank.s,
				Unknown: false,
			}

		}

	} else {
		fmt.Println(word, "1did hit 1000 !")
		fmt.Println(preresprank)
		return &Response{
			Word:    word,
			Rank:    preresprank.p,
			Score:   preresprank.s,
			Unknown: false,
		}

	}

	// fmt.Println("WTF")

	// if err := json.Unmarshal(resBody, &prerespnorank); err != nil {
	// 	// Word is not "recognised but outside first 1000"

	// 	if err = json.Unmarshal(resBody, &preresprank); err != nil {
	// 		// Word is not "recognised and in first 1000"

	// 		if err = json.Unmarshal(resBody, &preresperror); err != nil {
	// 			// Word is not "unrecognised"

	// 			log.Fatal(err)
	// 		} else {
	// 			fmt.Println(word, "is an unknown word for cemantix.")
	// 			resp.Unknown = true
	// 			return resp
	// 		}

	// 	} else {
	// 		fmt.Println(word, "did hit 1000 !")
	// 		resp.Rank = preresprank.p
	// 		resp.Score = preresprank.s
	// 		resp.Unknown = false
	// 		return resp
	// 	}

	// } else {
	// 	fmt.Println(word, "did not hit 1000 :", string(resBody))
	// 	resp.Rank = 0
	// 	resp.Score = prerespnorank.s
	// 	resp.Unknown = false
	// 	return resp
	// }
	return nil
}
