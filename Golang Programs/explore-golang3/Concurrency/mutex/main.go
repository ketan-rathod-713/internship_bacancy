package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var trackMutex sync.Mutex
var track [3]bool // lets say 3 lane route

type Utilization struct {
	TotalTime    time.Duration
	NumberOfCars int
}

var trackUtilization [3]Utilization

func main() {
	// var wg sync.WaitGroup
	// // for eg. i have 10 cars which will be using this tracks for 2 sec each.

	// for i := 1; i <= 10; i++ {
	// 	wg.Add(1)
	// 	go runningCar(i, &wg)
	// }

	// wg.Wait()

	//Alternate 2 Do above using channel

	var wg sync.WaitGroup
	var c chan int = make(chan int, 3) // let's say we have 3 tracks
	// initally 3 tracks ki information dal do
	c <- 0
	c <- 1
	c <- 2

	fmt.Println(trackUtilization)

	for i := 1; i < 6; i++ {
		wg.Add(1)
		go runCarChannel(i, c, &wg)
	}

	wg.Wait()
	fmt.Println(trackUtilization)
}

// jese hi koi track khali hoga usko ham channel me dal denge then koi bhi car usko access kar sakti he
func runCarChannel(id int, c chan int, wg *sync.WaitGroup) { // here channel refers to track id
	defer wg.Done()
	for i := 0; i < 2; i++ {
		trackId := <-c

		// run on this trackId
		startTime := time.Now()
		fmt.Printf("car %d is running on track %d \n", id, trackId)

		// run car on this track for 2 sec and then hat jao yaha se
		time.Sleep(time.Duration(rand.Intn(7) * int(time.Second)))
		endTime := time.Now()
		// calculate TrackUtilization
		timeOnTrack := endTime.Sub(startTime)

		currentTrackTime := trackUtilization[trackId].TotalTime
		trackUtilization[trackId].TotalTime = currentTrackTime + timeOnTrack
		trackUtilization[trackId].NumberOfCars += 1
		// vapas channel me track dal do ham side me khade he 3 seconds ke liye
		c <- trackId // jo track hamne liya tha vo vapas kar diya

		time.Sleep(time.Duration(rand.Intn(6) * int(time.Second)))
	}
}

func runningCar(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// fmt.Println("Car", id, "is acquiring lock for track")
	trackMutex.Lock()

	// check if any track if empty if empty then use it.
	trackToRun := -1
	for index, val := range track {
		if val == false {
			trackToRun = index
			break
		}
	}

	if trackToRun == -1 {
		fmt.Printf("car %d ko lock mil gaya but track nahi mila time par \n", id)
		trackMutex.Unlock()

		return
	}
	track[trackToRun] = true
	trackMutex.Unlock()

	fmt.Printf("Car %d has started running.. on track %d \n", id, trackToRun)
	// Now run here for 2 seconds and then remove from here.
	time.Sleep(2 * time.Second)

	trackMutex.Lock()

	track[trackToRun] = false

	trackMutex.Unlock()
	fmt.Printf("car %d has stopped running on track %d \n", id, trackToRun)
}
