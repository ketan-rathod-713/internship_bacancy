package channel

import "fmt"

func Channel() {

	messages := make(chan string)

	go func() { messages <- "Ping ha ha" }()

	msg := <-messages
	fmt.Println(msg)
}

func ChannelBuffering() {

	messages := make(chan string, 2) // at max 2 values can hold a channel

	go func() {
		messages <- "Ping ha ha"
		messages <- "This is second ping ha ha"
	}()

	fmt.Println(messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func ChannelSynchronization() {

	messages := make(chan string, 2) // at max 2 values can hold a channel

	go func() {
		messages <- "Ping ha ha"
		messages <- "This is second ping ha ha"
	}()

	fmt.Println(messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
