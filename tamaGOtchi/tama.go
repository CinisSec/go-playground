package tamagotchi

type Tama struct {
	// core
	health    int
	hunger    int
	happiness int
	energy    int
	alive     bool

	/* combat stats
	strength     int
	agility      int
	dexterity    int
	intelligence int
	luck         int
	*/
}

func newTama() *Tama {
	return &Tama{
		hunger:    100, // 0-100
		health:    100, // 0-100
		happiness: 50,  // 0-100
		energy:    0,
		alive:     true,

		/*
			strength:     1,
			agility:      1,
			dexterity:    1,
			intelligence: 1,
			luck:         1,
		*/
	}
}
