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

var itemCache = make(map[int]Item)

func init() {
	data, err := os.ReadFile("gamedata/items.json")
	if err != nil {
		log.Fatalf("Failed to read items.json: %v", err)
	}

	var items []Item
	if err := json.Unmarshal(data, &items); err != nil {
		log.Fatalf("Failed to parse items.json: %v", err)
	}

	for _, item := range items {
		itemCache[item.Id] = item
	}

	log.Printf("Loaded %d items from items.json", len(items))

}

func SearchItem(id int) Item {
	item, ok := itemCache[id]
	if !ok || item.Name == "" {
		return Item{}
	}
	return item
}

func LookupItemId(name string) int {
	for id, item := range itemCache {
		if item.Name == name {
			return id
		}
	}
	return 0
}

func ListItems() {
	for _, item := range itemCache {
		println(item.Name)
	}
}
