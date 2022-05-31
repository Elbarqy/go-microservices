package concrete

import (
	"designPatterns/creational/factory/interfaces"
)

type Musket struct {
	Gun
}

func NewMusket() interfaces.IGun {
	return &Musket{
		Gun{
			name:  "Musket",
			power: 10,
		},
	}
}
