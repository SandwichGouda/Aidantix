package main

import (
	"aidantix/db1"
	"fmt"
)

func main() {
	// req.Request("valeur")

	w := db1.ImportDatabase()

	for i, wd := range w {
		fmt.Println(i, wd.Label)
	}

	fmt.Println(w)
}
