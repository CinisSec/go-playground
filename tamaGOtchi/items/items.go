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

var itemCache []Item

func init() {
	data, err := os.ReadFile("gamedata/items.json")
	if err != nil {
		log.Fatalf("Failed to read items.json: %v", err)
	}

	var items []Item
	if err := json.Unmarshal(data, &items); err != nil {
		log.Fatalf("Failed to parse items.json: %v", err)
	}

	// Find max ID to size the map properly
	maxID := 0
	for _, item := range items {
		if item.Id > maxID {
			maxID = item.Id
		}
	}

	itemCache = make([]Item, maxID+1)
	for _, item := range items {
		itemCache[item.Id] = item
	}

	log.Printf("Loaded %d items from items.json", len(items))

}

func SearchItem(id int) Item {
	if id >= 0 && id < len(itemCache) && itemCache[id].Name != "" {
		return itemCache[id]
	}
	return Item{}
}

func LookupItemId(name string) (int, bool) {
	for id, item := range itemCache {
		if item.Name == name {
			return id, true
		}
	}
	return 0, false
}
