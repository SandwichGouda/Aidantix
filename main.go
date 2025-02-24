package main

import (
	"aidantix/readdb1"
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

func print_result() {
	//print result
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
			fmt.Print("Input convertion to int type failed. Please type in the puzzle day number : ")
			nb, _ = reader.ReadString('\n')
			nb = nb[:len(nb)-1]
			day, err = strconv.Atoi(nb)
		}
	}

	fmt.Println(`																 
																 
		 █████╗ ██╗██████╗  █████╗ ███╗   ██╗████████╗██╗██╗  ██╗
		██╔══██╗██║██╔══██╗██╔══██╗████╗  ██║╚══██╔══╝██║╚██╗██╔╝
		███████║██║██║  ██║███████║██╔██╗ ██║   ██║   ██║ ╚███╔╝ 
		██╔══██║██║██║  ██║██╔══██║██║╚██╗██║   ██║   ██║ ██╔██╗ 
		██║  ██║██║██████╔╝██║  ██║██║ ╚████║   ██║   ██║██╔╝ ██╗
		╚═╝  ╚═╝╚═╝╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═══╝   ╚═╝   ╚═╝╚═╝  ╚═╝
																 
																 `)

	fmt.Println("Cémantix : Jour", day)

	database := readdb1.ImportDatabase()

	var response *req.Response
	fmt.Println(nb_lines)
	for i, wd := range database {
		if i < nb_lines {
			continue
		}
		fmt.Println("Requesting :", wd.Label, "(", i, ")")
		response = req.Request(wd.Label, day)
		fmt.Println("Response :", *response)
		result.Write(response)
		time.Sleep(1000 * time.Millisecond)
	}
}

func false_main() {

	w := readdb1.ImportDatabase()

	for i, wd := range w {
		fmt.Println(i, wd.Label)
	}

	fmt.Println(w)
}

func false_main_2() {
	fmt.Println(result.InitCSV())

	r := &req.Response{
		Word:    "Cahak",
		Rank:    69,
		Score:   420.69,
		Unknown: false,
	}

	result.Write(r)
}
