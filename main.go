package main

import (
	"fmt"
	"github.com/gpbbit/gb_go_web/yaloader"
	"log"

	"github.com/gpbbit/gb_go_web/searcher"
)

func main() {
	fmt.Println("Lesson 1")
	sitesList := []string{
		"https://mebelvia.ru",
		"https://posudacenter.ru",
		"https://www.stolplit.ru",
		"https://lazurit.com",
		"http://blizzard.com/",
	}
	fmt.Printf("Got array %v\n", sitesList)
	result, err := searcher.SearchWordsOnPage("мебель", sitesList)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result array %v\n", result)
	link := "https://yadi.sk/i/xV3LGCEwax8RfA"
	fmt.Println("Download file ", link)
	path, err := yaloader.FileLoader(link)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File download to ", path)

}
