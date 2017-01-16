package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/Sirupsen/logrus"
	"github.com/l2x/golang-chinese-to-pinyin"
	"github.com/urfave/cli"
)

type Data struct {
	Name string `json:"name"`
	Img  string `json:"url"`
}

type Shicai struct {
	Cat  string
	URL  string
	Name string
}

var getImageCommand = cli.Command{
	Name:  "images",
	Usage: "get images of food data",
	Action: func(context *cli.Context) error {
		if context.NArg() != 1 {
			return fmt.Errorf("Usage: %s  %s <filename>", os.Args[0], context.Command.Name)
		}

		shicai, err := parseShicaiFile(context.Args()[0])
		if err != nil {
			logrus.Fatal(err)
		}

		// 获取图片下载地址
		for _, s := range shicai {
			doc, err := goquery.NewDocument(s.URL)
			if err != nil {
				logrus.Fatal(err)
			}
			doc.Find(".foopi").Each(func(i int, s *goquery.Selection) {
				url, _ := s.Find("img").Attr("src")
				if url != "" {
					fmt.Printf("%s\n", url)
				}
			})
		}
		return nil
	},
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

		for i, s := range shicai {
			fmt.Printf("%d: %s\n", i, s.Name)
			doc, err := goquery.NewDocument(s.URL)
			if err != nil {
				logrus.Fatal(err)
			}
			var d []Data
			doc.Find(".foopi").Each(func(i int, s *goquery.Selection) {
				name := s.Find("a").Text()
				url, _ := s.Find("img").Attr("src")
				t := Data{
					Name: name,
					Img:  url,
				}
				d = append(d, t)
			})
			saveData(d, s.Name+".json")
		}

		return nil
	},
}

var getCategoryCommand = cli.Command{
	Name:  "catagory",
	Usage: "get index of food data",
	Action: func(context *cli.Context) error {
		if context.NArg() != 1 {
			return fmt.Errorf("Usage: %s  %s <filename>", os.Args[0], context.Command.Name)
		}

		shicai, err := parseShicaiFile(context.Args()[0])
		if err != nil {
			logrus.Fatal(err)
		}

		for i, s := range shicai {
			fmt.Printf("%d: %s\n", i, s.Name)
			doc, err := goquery.NewDocument(s.URL)
			if err != nil {
				logrus.Fatal(err)
			}
			var d []string
			doc.Find(".foopi").Each(func(i int, s *goquery.Selection) {
				name := s.Find("a").Text()
				d = append(d, name)
			})
			saveCat(d, s.Name+".txt")
		}

		return nil
	},
}

func saveCat(d []string, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	tmp := strings.Join(d, "\n")
	_, err = io.WriteString(f, tmp)
	return err
}

func saveData(d []Data, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return writeJSON(f, d)
}

func writeJSON(w io.Writer, v interface{}) error {
	data, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

func parseShicaiFile(path string) ([]Shicai, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, err
	}

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
