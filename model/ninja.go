package model

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

type PriceData struct {
	Lines            []Item `json:"lines"`
	itemsKeyedByName map[string]*Item
}

type Item struct {
	ID                     int                 `json:"id"`
	Name                   string              `json:"name"`
	Icon                   string              `json:"icon"`
	BaseType               string              `json:"baseType"`
	StackSize              int                 `json:"stackSize"`
	ItemClass              int                 `json:"itemClass"`
	Sparkline              Sparkline           `json:"sparkline"`
	LowConfidenceSparkline Sparkline           `json:"lowConfidenceSparkline"`
	ImplicitModifiers      []any               `json:"implicitModifiers"`
	ExplicitModifiers      []ExplicitModifiers `json:"explicitModifiers"`
	FlavourText            string              `json:"flavourText"`
	ChaosValue             float64             `json:"chaosValue"`
	ExaltedValue           float64             `json:"exaltedValue"`
	DivineValue            float64             `json:"divineValue"`
	Count                  int                 `json:"count"`
	DetailsID              string              `json:"detailsId"`
	TradeInfo              []any               `json:"tradeInfo"`
	ListingCount           int                 `json:"listingCount"`
}

type Sparkline struct {
	Data        []float64 `json:"data"`
	TotalChange int       `json:"totalChange"`
}

type ExplicitModifiers struct {
	Text     string `json:"text"`
	Optional bool   `json:"optional"`
}

func FetchPriceData(url string, item_type string) PriceData {
	// "https://poe.ninja/api/data/itemoverview?league=Affliction&type="
	var data PriceData
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

func NewPriceData(url string, item_type string) *PriceData {
	// fetch data from poe.ninja
	priceData := FetchPriceData(url, item_type)
	// create map of items with names as keys
	itemsKeyedByName := make(map[string]*Item)
	for _, item := range priceData.Lines {
		itemsKeyedByName[item.Name] = &item
	}
	// return the poeninja data plus the map
	return &PriceData{
		Lines:            priceData.Lines,
		itemsKeyedByName: itemsKeyedByName,
	}
}

func (pd PriceData) GetItemChaosPrice(item_name string) float64 {
	for i := range pd.Lines {
		if pd.Lines[i].Name == item_name {
			return pd.Lines[i].ChaosValue
		}
	}
	return -1
}

func (pd PriceData) GetItemDivinePrice(item_name string) float64 {
	for i := range pd.Lines {
		if pd.Lines[i].Name == item_name {
			return pd.Lines[i].DivineValue
		}
	}
	return -1
}

func (pd PriceData) GetItemList() []string {
	var names []string
	for i := range pd.Lines {
		names = append(names, pd.Lines[i].Name)
	}
	return names
}

func (pd PriceData) SearchItems(term string) []string {
	return fuzzy.FindFold(term, pd.GetItemList())
}
