package goroutine

import (
	"fmt"
	"math/rand"
	"time"
)

func GoroutineCommunicationWithMain() {
	c := make(chan string)
	go boring("boring !!", c)

	for i := 0; i < 5; i++ { // if i increase value here then error
		fmt.Printf("You say %v \n", <-c) // receive expression is just a value
	}
}

func boring(msg string, c chan string) {
	for i := 0; i < 5; i++ { // fatal error: all goroutines are asleep - deadlock!
		c <- fmt.Sprintf("%s %d", msg, i) // expression to be send can be any suitable value
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func ConcurrencyPatternGenerator() {
	c := boring2("Boring!")
	secondChannel := boring2("Again Boring 2 !!")

	for i := 0; i < 5; i++ {
		fmt.Printf("%v \n", <-c)             // iski value isme de do ha ha means receiver type
		fmt.Printf("%v \n", <-secondChannel) // iski value isme de do ha ha means receiver type
	}

	fmt.Println("I am boring !! leaving now..")
}

func boring2(msg string) <-chan string { // returns recieve only channel of string TODO:
	c := make(chan string)
	go func() { // we launch goroutine from inside a function
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c // return the channel to the caller
}

/*
? Synchronization

When the main function executes <-c it will wait for a value to be sent
similarly when the boring function executes c <- value, it waits for reciever to be ready.
a sender and receiveer must both be ready to play their part in comuunication.

This channels both communicate and synchronize.


! Buffer channels
Buffering removes synchronizations
like mail box

? THE GO APPROACH
Dont communicate by sharing memory, share memory by communicating,
no need of mutexes  and all for shared memory
No need of locks

insteaad we can use channels to pass data go and forth.

? based on above principles : we got some concurrency patterns

! Generator pattern : Function that returns a channel

function that returns a channel
and inside it starts the concurrent computation inside it

!16:09 continue

? Channels are first class values just like string or integer.

? In this we launch goroutine inside the function itself and return channel itself so that we can get data computation concurrently

? It is just like the services, we can have multiple such services


TODO our boring function returns a channel that lets us communicate with the boring service it provides.

TODO we can have more instances of the service

! The problem here is that secondChannel has to wait for the first channel and hence it will be synchronised everytime. And hence this approach is not good for such cases. so lets see another case.

? Multiplexing : fan in function

? What we do here is we pass two channels and multiplex it to one channel, other two channels are now separate go routines inside that function for infinite time.


*/

// Here input 1 and input 2 can be at different levels depends
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1 // jo input one mese aaya he vo isme dal do
		}
	}()
	go func() {
		for {
			c <- <-input2 // jo input two mese aaya he vo isme dal do
		}
	}()

	return c // ab sab kuch c dekhega tum dono bas concurrently muje data bhejate rehna :)))
}

func FanInMultiplexingPattern() {
	c := fanIn(boring2("Jhon"), boring2("Eve"))
	for i := 0; i < 10; i++ {
		fmt.Printf("%v \n", <-c)
	}
}

// TODO: 18:29 Continue from here
// ? https://www.youtube.com/watch?v=f6kdp27TYZs

/*

What if we want them to be totally synchronous
Channels are first class values.
We can send channel and it can include channel inside it.

? Restoring Sequencing
? Send a channel on channel, making goroutine wait its turn.
? Recieve all messages, then enable them again by sending on private channel.
? First we define a message type that contains channel for the reply.

*/

type Message struct {
	str  string
	wait chan bool
}

// ! DONT UNDERSTOOD
func SynchronousFanInFunction() {
	c := FanIn3(boring3("Ketan"), boring3("Aman"))
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)

		// now mark them true so that it can procceed further
		msg1.wait <- true // we are sending information in this block
		msg2.wait <- true // same here we are sending information in this block hence when reciver accept i will go forward then only
	}
}

func boring3(message string) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool) // shared between all mesages.

	go func() {
		for i := 0; ; i++ {
			c <- Message{str: message, wait: waitForIt}
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			<-waitForIt // we are receiving information here hence blocking it
			// if not recieving information from this channal then will result in error as everything will be blocked and all that.
			// Hence we restored the sequencing here.
		}
	}()

	return c
}

func FanIn3(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)

	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}

/* By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables. */

/*
? Select

* A control structure unique to concurrency.
* The reason channels and goroutines are built into language.
* Just like switch = which communication we should proceed now. :)

? Provides another way to handle multiple channels.
it is like switch but here each case is communication.
all channels are evaluated.
selection blocks until one communication can procceed.
! if multiple can proceed, select statement chooses pseudo randomaly.
A default case :if presents, executes immediatly if no channel is ready.
*/

func FanInWithSelect(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	// written inside goroutine
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case p := <-input2:
				c <- p
			}
		}
	}()

	return c
}

func FanInWithSelectExample() {
	c := FanInWithSelect(boring2("Ketan selecct"), boring2("Aman select"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}
