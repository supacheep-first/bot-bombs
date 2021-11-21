package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func selectHero() bool {
	x := 1069
	y := 769
	_count := 10
	for _count > 0 {
		robotgo.Move(x, y) // hero
		color := robotgo.GetPixelColor(x, y)
		if color == "27283b" {
			robotgo.Click("left")
			fmt.Println("click hero")
			robotgo.Sleep(5)
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
	for i := 0; i < 3; i++ {
		robotgo.Move(432, 734)
		robotgo.MouseToggle("down")
		robotgo.MoveSmooth(427, 349, 1.0, 100.0)
		robotgo.MouseToggle("up")
		robotgo.Sleep(1)
	}

	// make hero go work NOW!!!
	for {
		x := 561
		y := 732
		robotgo.Move(x, y) // work
		color := robotgo.GetPixelColor(x, y)
		if color == "7c9e52" || color == "6f8d5a" || color == "73935e" || color == "739254" || color == "779858" || color == "57815b" {
			robotgo.Sleep(1)
			robotgo.Click("left")
			robotgo.Sleep(1)
		} else if color == "e6b7a7" {
			_over := 0
			for _over <= 3 {
				robotgo.Move(928, 371) // error overload
				robotgo.Sleep(1)
				robotgo.Click("left")
				robotgo.Sleep(1)
				fmt.Println("close overload")
				_over++
				break
			}
		} else if color == "a5d386" || color == "82c188" {
			break
		}
		robotgo.MilliSleep(20.0)
	}

	robotgo.Move(716, 355) // close dialog hero
	robotgo.Click("left")
	fmt.Println("click close dialog hero")
	robotgo.Sleep(10)
	return true
}
