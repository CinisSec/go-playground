package main

import (
	"tamaGOtchi/items"
)

func main() {
	//tama := creature.NewTama()
	//tama.PrintStats()

	items.LoadAllItems()
	println(items.SearchItem(100).Name)

}
