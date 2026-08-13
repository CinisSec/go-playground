package creature

type Tama struct {
	// core
	Health    int
	Hunger    int
	Happiness int
	Energy    int
	Alive     bool
}

func NewTama() *Tama {
	return &Tama{
		Hunger:    100, // 0-100
		Health:    100, // 0-100
		Happiness: 50,  // 0-100
		Energy:    0,
		Alive:     true,
	}
}

func (tama *Tama) PrintStats() {
	println("Stats: ")
	println("----------------")
	println("Health: ", tama.Health)
	println("Hunger: ", tama.Hunger)
	println("Happiness: ", tama.Happiness)
	println("Energy: ", tama.Energy)
	println("Alive: ", tama.Alive)
	println("----------------")
}
