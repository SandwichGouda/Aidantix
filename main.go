package main

import (
	"aidantix/db1"
	"fmt"
)

func main() {
	// req.Request("valeur")

	w := db1.ImportDatabase()

	fmt.Println(w)
}
