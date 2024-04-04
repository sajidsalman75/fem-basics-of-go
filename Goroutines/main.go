package main

import (
	"fmt"
	"time"
)

func printMessage(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Millisecond * 800)
	}
	channel <- "Done"
}
func main() {
	channel := make(chan string)
	go printMessage("Hello World!", channel)
	response := <-channel
	fmt.Println(response)
	// go printMessage("Frontend Masters Rocks!")
	// printMessage("Hello Another World!")
	// go printMessage("Hello Another World!")
}
