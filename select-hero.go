package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func selectHero() bool {
	x := 1042
	y := 769
	_count := 10
	for _count > 0 {
		robotgo.Move(x, y) // hero
		color := robotgo.GetPixelColor(x, y)
		if color == "2f2a37" {
			robotgo.Click("left")
			fmt.Println("click hero")
			break
		}
		robotgo.Sleep(1)
		_count--
	}
	// when waiting loading game long time #60s
	if _count <= 0 {
		doRefresh()
		return false
	}

	robotgo.Sleep(5)

	// move to bottom
	// for i := 0; i < 3; i++ {
	// 	robotgo.Move(432, 734)
	// 	robotgo.MouseToggle("down")
	// 	robotgo.MoveSmooth(427, 349, 1.0, 100.0)
	// 	robotgo.MouseToggle("up")
	// 	robotgo.Sleep(1)
	// }

	// make hero go work NOW!!!
	// for {
	// 	x := 574
	// 	y := 762
	// 	robotgo.Move(x, y) // work
	// 	color := robotgo.GetPixelColor(x, y)
	// 	if color == "7c9e52" || color == "6f8d5a" {
	// 		robotgo.Click("left")
	// 		robotgo.Sleep(1)
	// 	} else if color == "a5d386" {
	// 		break
	// 	}
	// }

	xx := 572
	yy := 380
	robotgo.Move(xx, yy) // work
	robotgo.Click("left")
	robotgo.Sleep(5)

	robotgo.Move(696, 334) // close dialog hero
	robotgo.Click("left")
	fmt.Println("click close dialog hero")
	robotgo.Sleep(5)
	return true
}
