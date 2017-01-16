package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/Sirupsen/logrus"
	"github.com/l2x/golang-chinese-to-pinyin"
	"github.com/urfave/cli"
)

type Shicai struct {
	Cat  string
	URL  string
	Name string
}

var getIndexCommand = cli.Command{
	Name:  "index",
	Usage: "get index of food data",
	Action: func(context *cli.Context) error {
		if context.NArg() != 1 {
			return fmt.Errorf("Usage: %s  %s <filename>", os.Args[0], context.Command.Name)
		}

		shicai, err := parseShicaiFile(context.Args()[0])
		if err != nil {
			logrus.Fatal(err)
		}

		// 获取图片下载地址
		//for _, s := range shicai {
		//	doc, err := goquery.NewDocument(s.URL)
		//	if err != nil {
		//		logrus.Fatal(err)
		//	}
		//	doc.Find(".foopi").Each(func(i int, s *goquery.Selection) {
		//		url, _ := s.Find("img").Attr("src")
		//		if url != "" {
		//			fmt.Printf("%s\n", url)
		//		}
		//	})
		//}

		for i, s := range shicai {
			fmt.Printf("%d: %s\n", i, s.Name)
			doc, err := goquery.NewDocument(s.URL)
			if err != nil {
				logrus.Fatal(err)
			}
			doc.Find(".foopi").Each(func(i int, s *goquery.Selection) {
				name := s.Find("a").Text()
				url, _ := s.Find("img").Attr("src")
			})
		}

		return nil
	},
}

func parseShicaiFile(path string) ([]Shicai, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return parseShicaiFromReader(f)
}

func parseShicaiFromReader(r io.Reader) ([]Shicai, error) {
	s := bufio.NewScanner(r)
	py := Pinyin.New()
	py.Split = "-"
	py.Upper = false
	var shicai []Shicai
	var t Shicai

	for s.Scan() {
		if err := s.Err(); err != nil {
			return nil, err
		}
		text := s.Text()
		parts := strings.SplitN(text, " ", 3)
		cat, _ := py.Convert(parts[0])
		t = Shicai{
			Cat:  cat,
			URL:  parts[1],
			Name: parts[2],
		}
		shicai = append(shicai, t)
	}

	return shicai, nil
}
