package result

import (
	"aidantix/req"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func CountLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lineCount, nil
}

// Output = i <=> Start parsing the database from i (there is a slight discrepancy : technically if the headers are there, it wont query for the first item in teh database.. But whatever really ("le" isn't that useful of a word))
func InitCSV() int {

	var (
		nb_lines int
		err      error
	)

	if !FileExists("result.csv") {
		file, err := os.Create("result.csv") // Creates or truncates an existing file
		if err != nil {
			log.Fatal("Error creating result.csv :", err)

		}
		file.Close()
	} else {
		nb_lines, err = CountLines("result.csv")
		if err != nil {
			fmt.Println("Failed to retrieve the number of lines in result.csv. Program will start from 0 (and append as if file was blank)")
		}
	}

	if nb_lines == 0 { // Add headers
		f, err := os.OpenFile("result.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println("Error opening result.csv :", err, ", carrying on")
		}
		defer f.Close()

		line := "Word,Rank,Score\n"

		if _, err := f.Write([]byte(line)); err != nil {
			f.Close()
			fmt.Println("Error appending headers to result.csv :", err, ", carrying on")
		}
	}

	return nb_lines
}

func Write(resp *req.Response) {

	f, err := os.OpenFile("result.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	
	if err != nil {
		log.Fatal("Error opening result.csv :", err)
	}
	defer f.Close()

	line := resp.Word + "," + strconv.Itoa(resp.Rank) + "," + strconv.FormatFloat(resp.Score, 'f', 4, 64) + " " + "\n"

	if _, err := f.Write([]byte(line)); err != nil {
		f.Close()
		log.Fatal("Error appending data to result.csv :", err)
	}

}
