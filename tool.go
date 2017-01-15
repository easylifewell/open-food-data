package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <url>", os.Args[0])
	}
	doc, err := goquery.NewDocument(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".foopi").Each(func(i int, s *goquery.Selection) {
		name := s.Find("a").Text()
		//		url, _ := s.Find("img").Attr("src")
		//		fmt.Printf("%d: %s(%s)\n", i, name, url)

		fmt.Printf("%s\n", name)

	})
}
