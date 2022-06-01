package buildMethods

import (
	"designPatterns/creational/builder/builders"
)

type IBuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() builders.House
}

func GetBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return &builders.NormalBuilder{}
	}
	return nil
}
