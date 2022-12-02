package main

import "fmt"

func main() {
	//Array and loops
	loops := 1
	var x [10]int
	y := [10]int{1, 21, 35, 4, 30, 6, 12, 9, 4, 10}
	for i := 0; i < len(y); i++ {
		fmt.Println(y[i])
	}
	for i, v := range y {
		fmt.Println(i, v)
	}
	x[5] = 100
	fmt.Println(x)
	for loops < 10 {
		fmt.Println(loops)
		loops++
	}

	for i := 0; i < loops; i++ {
		fmt.Println(i)
	}

	//slices similar to vectors

	// strSlice := make([]string, 10)
	var creatures = [4]string{"Pokemon", "Furby", "Unicorn", "Gopher"}

	slice1 := creatures[:2]
	slice2 := creatures[1:2]
	slice3 := creatures[:3]
	fmt.Println(slice1, slice2, slice3)

	//Maps
	strMap := make(map[string]map[string]string)

	strMap = map[string]map[string]string{
		"P": {
			"name":  "Pokemon",
			"power": "lighting",
		},
	}
	strMap["s"] = map[string]string{
		"name":  "seer",
		"power": "Prediction",
	}
	fmt.Println(strMap)
	delete(strMap, "s")
	fmt.Println("after deletion", strMap)
	// map iteration
}
