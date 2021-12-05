package main

import "github.com/go-vgo/robotgo"

func openTabGame() {
	robotgo.Move(921, 212) // paly now
	robotgo.Click("left")
	robotgo.Sleep(60)
}
