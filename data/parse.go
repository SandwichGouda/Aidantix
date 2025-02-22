package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func umain() {
	// Read JSON file
	data, err := ioutil.ReadFile("mots2.json")
	if err != nil {
		log.Fatal(err)
	}

	// Define a variable to hold the JSON data
	var jsonArray []map[string]interface{}

	// Unmarshal JSON into the slice
	err = json.Unmarshal(data, &jsonArray)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate over the array
	for i, item := range jsonArray {
		fmt.Println(i, item["label"])
	}
}
