package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(350*time.Millisecond)
	boom := time.After(5*time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("tick...")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default: // it runs if no cases are ready
			fmt.Println("    ...")
			time.Sleep(time.Second)
		}
	}
}