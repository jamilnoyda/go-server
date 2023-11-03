package main

import (
	"fmt"

	mymath1 "github.com/jamilnoyda/go-server/math"
	ping "github.com/go-ping/ping"
	"time"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(mymath1.Add(2, 2))

	pinger, err := ping.NewPinger("google.com")
	if err != nil {
		logger.Error(err)
		return false
	}
	pinger.Count = 1
	pinger.Timeout = 5 * time.Second
	err = pinger.Run() // Blocks until
	if err != nil {
		logger.Error("Unable to Ping the IP:", err)
		time.Sleep(5 * time.Second)
	}

}
