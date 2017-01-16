package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/urfave/cli"
)

var getIndexCommand = cli.Command{
	Name:  "index",
	Usage: "get index of food data",
	Action: func(context *cli.Context) error {
		if context.NArg() != 1 {
			return fmt.Errorf("Usage: %s  %s url", os.Args[0], context.Command.Name)
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

		return nil
	},
}
