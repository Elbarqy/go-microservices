package directors

import (
	builderMethods "designPatterns/creational/builder/builderMethods"
	"designPatterns/creational/builder/builders"
)

type Director struct {
	builder builderMethods.IBuilder
}

func NewDirector(b builderMethods.IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b builderMethods.IBuilder) {
	d.builder = b
}

func (d *Director) BuildHouse() builders.House {
	d.builder.SetDoorType()
	d.builder.SetWindowType()
	d.builder.SetNumFloor()
	return d.builder.GetHouse()
}
