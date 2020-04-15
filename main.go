package main

import (
	"flag"
	"fmt"
	"github.com/gpbbit/gb_go_web/searcher"
	"github.com/gpbbit/gb_go_web/yaloader"
	"log"
	"os"
)

func main() {
	var lessonNum int
	var taskNum int
	flag.IntVar(&lessonNum, "lesson", 1, "lesson number")
	flag.IntVar(&taskNum, "task", 1, "task number")
	flag.Parse()

	switch lessonNum {
	case 1:
		fmt.Println("Lesson 1")
		switch taskNum {
		case 1:
			sitesList := make([]string, 0, 15)
			var query string
		Loop:
			for {
				var link string
				fmt.Print("Input link for search: ")
				fmt.Fscan(os.Stdin, &link)
				fmt.Println(link)
				switch link {
				case "\n":
					fmt.Println("Empty data, is very bad!")
					break
				case ":q":
					fmt.Println("Good!")
					break Loop
				default:
					sitesList = append(sitesList, link)
				}
			}
			fmt.Printf("Got array %v\n", sitesList)
			fmt.Print("Input query data for search: ")
			fmt.Fscan(os.Stdin, &query)
			result, err := searcher.SearchWordsOnPage(query, sitesList)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Result array %v\n", result)
		case 2:
			fmt.Println("Download file from yadi.sk")
			var link string
			for {
				fmt.Print("Input link for download: ")
				fmt.Fscan(os.Stdin, &link)
				if link != "" && link != " " {
					break
				} else {
					fmt.Println("Empty link, is very bad!")
				}
			}
			//link := "https://yadi.sk/i/xV3LGCEwax8RfA"
			fmt.Println("Download file ", link)
			path, err := yaloader.FileLoader(link)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("File download to ", path)
		default:
			fmt.Println("Task not found")
		}
	default:
		fmt.Println("Lesson not found")
	}

}
