package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

// func main() {
// 	get()
// }

func get() {
	robotgo.MilliSleep(2)
	x, y := robotgo.GetMousePos()
	fmt.Println("pos: ", x, y)

	color := robotgo.GetPixelColor(x, y)
	fmt.Println("color---- ", color)
}

//789f39
//7cb27b
