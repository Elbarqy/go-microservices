package main

import (
	factories "designPatterns/creational/factory/factory"
	"fmt"
)

func main() {
	musket, err := factories.GunFactory("musket")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(musket.GetName())
}
