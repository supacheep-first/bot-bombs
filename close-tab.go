package main

import "github.com/go-vgo/robotgo"

func closeTab() {
	robotgo.Move(477, 77) // close tab game
	robotgo.Click("left")
	robotgo.Sleep(10)
}
