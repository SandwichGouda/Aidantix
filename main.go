package main

import (
	"aidantix/req"
	"fmt"
)

type Response struct {
	Word    string
	Rank    int
	Score   float64
	Unknown bool
}

// func (resp Response) Format() string {
// 	if tmprtr_str, err := strconv.ParseFloat(resp.Temperature, 64); err == nil {
// 		if resp.Rank == 0 {
// 			return resp.Word + " - " + tmprtr_str
// 		}
// 		return resp.Word + " ! " + string(resp.Rank) + " - " + tmprtr_str
// 	} else {
// 		log.Fatal(err)
// 	}

// }

func main() {

	fmt.Println(req.Request("valeur"))

	// w := db1.ImportDatabase()

	// for i, wd := range w {
	// 	fmt.Println(i, wd.Label)
	// }

	// fmt.Println(w)

	// db1.Write()
}
