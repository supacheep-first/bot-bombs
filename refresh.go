package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func doRefresh() {
	robotgo.Sleep(3)
	robotgo.KeyDown("lctrl")
	robotgo.KeyTap("f5")
	robotgo.KeyUp("lctrl")
	fmt.Println("refresh page.")
}
