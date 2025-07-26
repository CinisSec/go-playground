package items

import (
	"encoding/json"
	"log"
	"os"
)

type Item struct {
	Id              int
	Name            string
	Description     string
	EconomicalValue int
	InInventory     int
	Consumable      bool
}

func LoadAllItems() []Item {
	var items []Item
	data, err := os.ReadFile("gamedata/items.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &items)
	if err != nil {
		log.Fatal(err)
	}

	// Make sure the index of the array is the same as item Id for quick search.
	itemMap := make([]Item, items[len(items)-1].Id+1)
	for _, item := range items {
		itemMap[item.Id] = item
	}
	return itemMap
}

func SearchItem(id int) Item {
	itemMap := LoadAllItems()
	if id >= 0 && id < len(itemMap) {
		return itemMap[id]
	}
	return Item{}
}
