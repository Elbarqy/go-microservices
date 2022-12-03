package patterns

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 2; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func BasicRoutineRunner() {
	defer fmt.Println("done")
	go say("hello")
	say("world") //when commenting this line the function doesn't block and return only done sometimes one run of the previous function
}
