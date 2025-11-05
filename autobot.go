package main

import "github.com/go-vgo/robotgo"

func main() {
	for {
		x, y := robotgo.Location()
		robotgo.Move(x+100, y+100)
		robotgo.Move(x, y)
		robotgo.Sleep(60)
	}
}
