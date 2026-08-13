package main

import (
	"tamaGOtchi/creature"
	"tamaGOtchi/items"
)

func main() {
	tama := creature.NewTama()
	tama.PrintStats()

	items.ListInventory()

	item := items.SearchItem(100)
	println(item.Name)

	item = items.SearchItem(2)
	println(item.Name)

	itemid := items.LookupItemId("Book")
	println(itemid)

	item = items.SearchItem(2)
	println(item.Name)
	println(item.Description)
	println(item.EconomicalValue)
	println(item.Consumable)

}
