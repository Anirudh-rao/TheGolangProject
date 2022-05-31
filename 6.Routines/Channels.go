package main

//Channels are the pipes that connect concurrent goroutines.
import (
	"fmt"
)

func main() {
	//To Make Chanel
	messages := make(chan string)
	// <- operator sends message to the channel
	go func() {
		messages <- "ping"
	}()
	// <- recieves the value from channel outsitde of Routine
	msg := <-messages
	fmt.Println(msg)

}
