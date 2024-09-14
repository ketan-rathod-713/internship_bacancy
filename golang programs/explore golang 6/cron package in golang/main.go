package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()

	c.AddFunc("* * * * 1", func() {
		fmt.Println("Every 1 minutes")
	})

	c.AddFunc("@every 1s", func() {
		fmt.Println("every 1 second")
	})

	fmt.Println(c.Entries())

	c.Start()

	select {}
}
