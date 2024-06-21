# Worker Pattern

The opposite pattern to fan-in is a fan-out or workers pattern. 
Multiple goroutines can read from a single channel, distributing an amount of work between CPU cores, hence the workers name. In Go, this pattern is easy to implement - just start a number of goroutines with channel as parameter, and just send values to that channel - distributing and multiplexing will be done by Go runtime, automagically :)