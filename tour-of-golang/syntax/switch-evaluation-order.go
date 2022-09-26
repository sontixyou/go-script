package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("What's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far way.")
	}
}

// 時間の足し算と減算はRubyみたいに直感ではいかない。
