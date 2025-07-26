package main

import (
	"tamaGOtchi/creature"
	"tamaGOtchi/items"
)

func main() {
	tama := creature.NewTama()
	tama.PrintStats()

	item := items.SearchItem(100)
	println(item.Name)

	item = items.SearchItem(2)
	println(item.Name)

	itemid, _ := items.LookupItemId("Book")
	println(itemid)

	items.ListInventory()

}
