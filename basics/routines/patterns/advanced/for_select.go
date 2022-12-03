package advanced

import "fmt"

func ForSelector() {
	done := make(chan interface{})
	// <- output stream
	chanOwner := func() <-chan string {
		stringStream := make(chan string)
		go func() {
			defer close(stringStream)
			for _, s := range []string{"a", "b", "c"} {
				select {
				case <-done:
					break
				case stringStream <- s:
				}
			}
		}()
		return stringStream
	}
	consumer := func(results <-chan string) {
		for result := range results {
			fmt.Println(result)
			if result == "b" {
				done <- " "
			}
		}
	}
	results := chanOwner()
	consumer(results)
}
