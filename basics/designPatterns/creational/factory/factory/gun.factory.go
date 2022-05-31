package factories

import (
	"designPatterns/creational/factory/concrete"
	"designPatterns/creational/factory/interfaces"
	"fmt"
)

func GunFactory(gunType string) (interfaces.IGun, error) {
	if gunType == "musket" {
		return concrete.NewMusket(), nil
	}
	return nil, fmt.Errorf("wrong_gun_type")
}
