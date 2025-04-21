package main

import (
	"aidantix/readdb1"
	"aidantix/readdb2"
	"aidantix/req"
	"aidantix/result"
	"aidantix/scrape"
	"aidantix/scrape_goquery"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Response struct {
	Word    string
	Rank    int
	Score   float64
	Unknown bool
}

func main() {

	var day int

	nb_lines := result.InitCSV()

	day = scrape.ScrapeDay()
	if day == 0 {
		fmt.Println("scrape.ScrapeDay() returned 0. Moving onto scrape_goquery.ScrapeDay()")
	}

	day = scrape_goquery.ScrapeDay()

	if day == 0 {
		fmt.Print("scrape_goquery.ScrapeDay() returned 0. Please type in the puzzle day number : ")
		reader := bufio.NewReader(os.Stdin)
		nb, _ := reader.ReadString('\n')
		nb = nb[:len(nb)-1]

		var err error
		day, err = strconv.Atoi(nb)

		for err != nil {
			fmt.Print("Input conversion to int type failed. Please type in the puzzle day number : ")
			nb, _ = reader.ReadString('\n')
			nb = nb[:len(nb)-1]
			day, err = strconv.Atoi(nb)
		}
	}

	fmt.Print("How much time should we wait between each request (in milliseconds) ? ")
	reader := bufio.NewReader(os.Stdin)
	buff_str, _ := reader.ReadString('\n')
	buff_str = buff_str[:len(buff_str)-1]

	buff, err := strconv.Atoi(buff_str)
	for err != nil {
		fmt.Println("Input conversion to int type failed.")
		fmt.Println("How much time should we wait between each request (in milliseconds) ? ")
		buff_str, _ = reader.ReadString('\n')
		buff_str = buff_str[:len(buff_str)-1]
		buff, err = strconv.Atoi(buff_str)
	}

	fmt.Print("Run in verbose mode ? (y/N)")
	reader = bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	answer = answer[:len(answer)-1]
	for answer != "n" && answer != "N" && answer != "y" && answer != "Y" && answer != "" {
		fmt.Print("Please type either 'y' or 'n'.")
		answer, _ = reader.ReadString('\n')
		answer = answer[:len(answer)-1]
	}

	verbose := answer == "y" || answer == "Y"

	fmt.Println(`																 
																 
		 █████╗ ██╗██████╗  █████╗ ███╗   ██╗████████╗██╗██╗  ██╗
		██╔══██╗██║██╔══██╗██╔══██╗████╗  ██║╚══██╔══╝██║╚██╗██╔╝
		███████║██║██║  ██║███████║██╔██╗ ██║   ██║   ██║ ╚███╔╝ 
		██╔══██║██║██║  ██║██╔══██║██║╚██╗██║   ██║   ██║ ██╔██╗ 
		██║  ██║██║██████╔╝██║  ██║██║ ╚████║   ██║   ██║██╔╝ ██╗
		╚═╝  ╚═╝╚═╝╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═══╝   ╚═╝   ╚═╝╚═╝  ╚═╝
																 
																 `)

	fmt.Println("Cémantix : Jour", day)

	var response *req.Response
	if verbose {
		fmt.Printf("%d lines found in result.csv. Reading database from index %d.\n", nb_lines, nb_lines)
	}

	database1 := readdb2.ImportDatabase()

	for i, wd := range database1 {
		if i < nb_lines {
			continue
		}
		if verbose {
			fmt.Printf("Requesting : %s (%d)\n", wd, i)
		}
		response = req.Request(wd, day)
		if verbose {
			fmt.Println("Response :", *response)
		}
		if !response.Unknown && response.Rank > 0 {
			result.Write(response)
		}
		time.Sleep(time.Duration(buff) * time.Millisecond)
	}

	/*
	database2 := readdb2.ImportDatabase()

	for i, wd := range database2 {
		if i < nb_lines {
			continue
		}
		if verbose {
			fmt.Println("Requesting :",wd,"(",i)
		}
		response = req.Request(wd, day)
		if verbose {
			fmt.Println("Response :", *response)
		}
		if !response.Unknown && response.Rank > 0 {
			result.Write(response)
		}
		time.Sleep(time.Duration(buff) * time.Millisecond)
	}
	*/

	fmt.Println("Done.")
}