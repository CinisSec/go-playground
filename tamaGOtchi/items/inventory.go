package items

import (
	"encoding/json"
	"log"
	"os"
)

type InventoryItem struct {
	Id          int
	InInventory int
}

type Inventory struct {
	Items []InventoryItem
}

var inventoryCache []InventoryItem

func init() {
	data, err := os.ReadFile("gamedata/inventory.json")
	if err != nil {
		log.Fatalf("Failed to read inventory.json: %v", err)
	}

	var inventoryitems []InventoryItem
	if err := json.Unmarshal(data, &inventoryitems); err != nil {
		log.Fatalf("Failed to parse items.json: %v", err)
	}

	// Find max ID to size the map properly
	maxID := 0
	for _, inventoryitem := range inventoryitems {
		if inventoryitem.Id > maxID {
			maxID = inventoryitem.Id
		}
	}

	inventoryCache = make([]InventoryItem, maxID+1)
	for _, inventoryitem := range inventoryitems {
		inventoryCache[inventoryitem.Id] = inventoryitem
	}

	log.Printf("Loaded %d items from inventory.json", len(inventoryitems))

}

func GetInventoryItem(id int) (InventoryItem, bool) {
	if id >= 0 && id < len(inventoryCache) && inventoryCache[id].Id != 0 {
		return inventoryCache[id], true
	}
	return InventoryItem{}, false
}

func ListInventory() bool {
	if len(inventoryCache) != 0 {
		for _, item := range inventoryCache {
			var itemname = SearchItem(item.Id).Name
			if itemname != "" {
				println("Item: ", itemname)
				println("Quantity: ", item.InInventory)
			}
		}
		return true
	}
	println("Inventory is empty")
	return false
}
