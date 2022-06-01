package main

import (
	buildMethods "designPatterns/creational/builder/builderMethods"
	directors "designPatterns/creational/builder/director"
	"fmt"
)

func main() {

	normalBuilder := buildMethods.GetBuilder("normal")

	director := directors.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()
	fmt.Printf("Normal House Door Type: %s\n", normalHouse.DoorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.WindowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.Floor)
}
