package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func selectHero2() bool {
	fmt.Printf("-------- start select hero")
	x := 1069
	y := 769
	_count := 10
	for _count > 0 {
		robotgo.Move(x, y) // hero
		color := robotgo.GetPixelColor(x, y)
		if color == "27283b" {
			robotgo.Click("left")
			fmt.Println("click hero")
			break
		}

		robotgo.Sleep(1)
		_count--
	}
	// when waiting loading game long time #60s
	if _count <= 0 {
		fmt.Errorf("error loading long time")
		return false
	}

	robotgo.Sleep(5)

	// move to bottom
	for i := 0; i < 2; i++ {
		robotgo.Move(432, 734)
		robotgo.MouseToggle("down")
		robotgo.MoveSmooth(427, 349, 1.0, 100.0)
		robotgo.MouseToggle("up")
		robotgo.Sleep(1)
	}

	// make hero go work NOW!!!
	loop := 4
	for loop > 0 {
		stop := false
		// move mouse to space
		robotgo.Move(1028, 518)

		button_X := 511
		button_Y := 768

		top_X := 511
		top_Y := 416

		for button_Y >= top_Y {
			_enegyColor := robotgo.GetPixelColor(button_X, button_Y)
			// fmt.Println("point: ", button_X, button_Y, _enegyColor)
			if _enegyColor == "b1e276" {
				// found energy for run!!!
				fmt.Println("found energy for run!!!")
				_buttonWorkColor := robotgo.GetPixelColor(button_X+40, button_Y)
				fmt.Println("point: ", button_X+40, button_Y, _buttonWorkColor)
				if _buttonWorkColor == "6f8d5a" || _buttonWorkColor == "7c9e52" || _buttonWorkColor == "73935e" {
					fmt.Println("move mouse: ", button_X+40, button_Y)
					robotgo.Move(button_X+40, button_Y)
					robotgo.Sleep(1)
					fmt.Println("click")
					robotgo.Click()
					robotgo.Sleep(1)

					_overloadColor := robotgo.GetPixelColor(button_X+40, button_Y)
					if _overloadColor == "e6b7a7" {
						robotgo.Move(928, 371) // error overload
						robotgo.Sleep(1)
						robotgo.Click("left")
						robotgo.Sleep(1)
						fmt.Println("close overload")
					}

					// move mouse to space
					fmt.Println("move mouse to space")
					robotgo.Move(1028, 518)
					button_Y++

				} else if _buttonWorkColor == "a5d386" {
					stop = true
					break
				}
			}

			button_Y--
			robotgo.MilliSleep(50.0)
		}

		if stop {
			break
		}

		robotgo.Move(top_X, top_Y)
		robotgo.MouseToggle("down")
		robotgo.MoveSmooth(top_X, 580, 1.0, 100.0)
		robotgo.MouseToggle("up")
		robotgo.Sleep(1)
		loop--
	}

	robotgo.Move(716, 355) // close dialog hero
	robotgo.Click("left")
	fmt.Println("click close dialog hero")
	robotgo.Sleep(3)
	fmt.Printf("-------- end select hero")
	return true
}
