package main

import (
	"tamaGOtchi/creature"
	"tamaGOtchi/items"
)

func main() {
	tama := creature.NewTama()
	tama.PrintStats()

	item, _ := items.SearchItem(100)
	println(item.Name)

	item, _ = items.SearchItem(2)
	println(item.Name)

	itemid, _ := items.LookupItemId("Book")
	println(itemid)

}
