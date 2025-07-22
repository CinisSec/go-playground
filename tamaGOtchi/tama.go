package main

type Tama struct {
	// core
	health    int
	hunger    int
	happiness int
	energy    int
	alive     bool
}

func newTama() *Tama {
	return &Tama{
		hunger:    100, // 0-100
		health:    100, // 0-100
		happiness: 50,  // 0-100
		energy:    0,
		alive:     true,
	}
}

func (tama *Tama) printStats() {
	println("Stats: ")
	println("----------------")
	println("Health: ", tama.health)
	println("Hunger: ", tama.hunger)
	println("Happiness: ", tama.happiness)
	println("Energy: ", tama.energy)
	println("Alive: ", tama.alive)
	println("----------------")
}
