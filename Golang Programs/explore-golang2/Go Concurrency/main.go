package main

import (
	"concurrency/channel"
	"concurrency/goroutine"
)

func main() {
	channel.Channel()

	channel.ChannelBuffering()

	// goroutine.GoroutineCommunicationWithMain()

	// goroutine.ConcurrencyPatternGenerator()

	// goroutine.FanInMultiplexingPattern()

	// goroutine.SynchronousFanInFunction()

	goroutine.FanInWithSelectExample()

	goroutine.TimeoutMain()

	goroutine.TimeoutWholeConversionUsingSelectMain()
}
