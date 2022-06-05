package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Trade struct {
	TradeDate   string `json:"date"`
	GroupId     string `json:"group_id"`
	Leg         string `json:"leg"`
	Currency    string `json:"currency"`
	Amount      int    `json:"amount"`
	Fee         int    `json:"fee"`
	Exponent    int    `json:"exponent"`
	Description string `json:"description"`
	Direction   string `json:"direction"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	jsonFile, err := os.Open("sample.json")
	check(err)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var trades []Trade
	json.Unmarshal(byteValue, &trades)
	fmt.Println(trades)
}
