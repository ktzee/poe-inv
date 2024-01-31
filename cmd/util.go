package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/ktzee/poe-inv/model"
)

func fetchPriceData(url string, item_type string) *model.PriceData {
	// "https://poe.ninja/api/data/itemoverview?league=Affliction&type="
	data := new(model.PriceData)
	resp, err := http.Get(url + item_type)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, &data)
	return data
}

func getItemChaosPrice(item_name string, pd model.PriceData) float64 {
	for i := range pd.Lines {
		if pd.Lines[i].Name == item_name {
			return pd.Lines[i].ChaosValue
		}
	}
	return -1
}

func getItemDivinePrice(item_name string, pd model.PriceData) float64 {
	for i := range pd.Lines {
		if pd.Lines[i].Name == item_name {
			return pd.Lines[i].DivineValue
		}
	}
	return -1
}
