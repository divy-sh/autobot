package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	sleep := flag.Int("seconds", 60, "time interval in seconds between each sleep timer reset")
	flag.Parse()

	*sleep = max(1, min(*sleep, 60*60)) // minimum sleep time of 1 second, maximum of 1 hour

	fmt.Printf(`
	   █████████  █████  ████████████████  ███████   ███████████     ███████    ███████████
  	  ███░░░░░███░░███  ░░███░█░░░███░░░████░░░░░███░░███░░░░░███  ███░░░░░███ ░█░░░███░░░█
 	 ░███    ░███ ░███   ░███░   ░███  ░███     ░░███░███    ░███ ███     ░░███░   ░███  ░ 
 	 ░███████████ ░███   ░███    ░███  ░███      ░███░██████████ ░███      ░███    ░███    
 	 ░███░░░░░███ ░███   ░███    ░███  ░███      ░███░███░░░░░███░███      ░███    ░███    
 	 ░███    ░███ ░███   ░███    ░███  ░░███     ███ ░███    ░███░░███     ███     ░███    
 	 █████   █████░░████████     █████  ░░░███████░  ███████████  ░░░███████░      █████   
	 ░░░░░   ░░░░░  ░░░░░░░░     ░░░░░     ░░░░░░░   ░░░░░░░░░░░     ░░░░░░░       ░░░░░    
	
	
	                 Keeping your device awake every %d seconds.
	
	 `, *sleep)

	for {
		x, y := robotgo.Location()
		now := time.Now().Format(time.TimeOnly)
		fmt.Printf("\rLast run: %s", now)
		robotgo.Move(x+100, y+100)
		robotgo.Move(x, y)
		time.Sleep(time.Duration(*sleep) * time.Second)
	}
}
