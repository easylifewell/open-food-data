package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/urfave/cli"
)

var getFoodDataCommand = cli.Command{
	Name:  "get",
	Usage: "get food data form douguo",
	Action: func(context *cli.Context) error {
		if context.NArg() != 1 {
			return fmt.Errorf("Usage: %s  %s <food name>", os.Args[0], context.Command.Name)
		}

		var food Food
		foodName := context.Args()[0]
		food.Name = foodName
		food.Jieshao = GetFoodData(foodName, "jieshao")
		food.Gongxiao = GetFoodData(foodName, "gongxiao")
		food.FoodValue = GetFoodDataValue(foodName)
		food.Yingyangjiazhi = GetFoodData(foodName, "yingyangjiazhi")
		food.Shiyongxiaoguo = GetFoodData(foodName, "shiyongxiaoguo")
		food.Shiyongrenqun = GetFoodData(foodName, "shiyongrenqun")
		food.howToSelect = GetFoodData(foodName, "howtoselect")
		food.howToStorage = GetFoodData(foodName, "howtostorage")

		res, _ := json.MarshalIndent(food, "", "\t")

		fmt.Println(string(res))
		return nil
	},
}

type FoodData struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ValueData struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type FoodDataValue struct {
	Title string      `json:"title"`
	Data  []ValueData `json:"data"`
}

type Food struct {
	Name           string        `json:"name"`
	Jieshao        FoodData      `json:"jieshao"`
	FoodValue      FoodDataValue `json:"foodvalue"`
	Gongxiao       FoodData      `json:"gongxiao"`
	Yingyangjiazhi FoodData      `json:"yingyangjiazhi"`
	Shiyongxiaoguo FoodData      `json:"shiyongxiaoguo"`
	Shiyongrenqun  FoodData      `json:"shiyongrenqun"`
	howToSelect    FoodData      `json:"howtoselect"`
	howToStorage   FoodData      `json:"howtostorage"`
}

func GetFoodDataValue(foodName string) FoodDataValue {
	var res FoodDataValue

	url := fmt.Sprintf("http://www.douguo.com/ingredients/%s/foodvalue", foodName)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	// Find the title
	res.Title = doc.Find("div.bkmcot").Find("h3").Text()
	// Find the content
	doc.Find("div.bkmcot").Find("tr").Each(func(i int, s *goquery.Selection) {
		var value ValueData
		value.Name = s.Find("td").Eq(0).Text()
		value.Value = s.Find("td").Eq(1).Text()
		res.Data = append(res.Data, value)
		value.Name = s.Find("td").Eq(2).Text()
		value.Value = s.Find("td").Eq(3).Text()
		res.Data = append(res.Data, value)
	})

	return res
}
func GetFoodData(foodName string, cat string) FoodData {
	var res FoodData
	url := fmt.Sprintf("http://www.douguo.com/ingredients/%s/%s", foodName, cat)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	var des []string
	// Find the title
	res.Title = doc.Find("div.bkmcot").Find("h3").Text()
	// Find the content
	doc.Find("div.bkmcot").Find("p").Each(func(i int, s *goquery.Selection) {
		des = append(des, s.Text())
	})

	res.Content = strings.Join(des, "\n")

	return res
}
