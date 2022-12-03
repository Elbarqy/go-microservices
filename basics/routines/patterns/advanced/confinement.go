package advanced

import "fmt"

//ensuring information is only ever available from one concurrent process

// TWO TYPES ad hoc confinement and lexical confinement

// Ad hoc confinement
// confinement through convention
func AdHocConfinement() {
	data := []int{1, 2, 3, 4}
	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}
	handleData := make(chan int)
	//pass data to channel then close chan
	go loopData(handleData)
	//loop through channel data
	for num := range handleData {
		fmt.Println(num)
	}
}

// lexical confinement
// as per it's name suggests / limit the variables within a scope
func LexicalConfinement() {
	chanOwner := func() <-chan int {
		result := make(chan int, 5)
		//generating values
		go func() {
			defer close(result)
			for i := 0; i <= 5; i++ {
				result <- i
			}
		}()
		return result
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done recieving!")
	}
	results := chanOwner()
	consumer(results)
}
