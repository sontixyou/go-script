package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// switch true　と書くのと一緒である。
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning.")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
