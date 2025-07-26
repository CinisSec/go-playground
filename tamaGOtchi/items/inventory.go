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

func LoadInventory() []InventoryItem {
	var inventory []InventoryItem
	data, err := os.ReadFile("gamedata/inventory.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &inventory)
	if err != nil {
		log.Fatal(err)
	}
	return inventory
}
