package patterns

import (
	"fmt"
	"sync"
	"time"
)

func countDown(n int) {
	for n >= 0 {
		fmt.Println(n)
		n--
		time.Sleep(time.Millisecond * 500)
	}
}

func MainCounter() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		countDown(3)
		wg.Done() //remove one from the counter
	}()
	wg.Add(1)
	go func() {
		countDown(3)
		wg.Done() //remove one from the counter
	}()
	go func() {
		countDown(100)
		wg.Done() //remove one from the counter
	}()
	wg.Wait()
}
