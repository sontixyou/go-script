package main

import "fmt"

func main() {
	defer fmt.Println("22222")
	defer fmt.Println("world")
	defer fmt.Println("111111")

	fmt.Println("Hello")
}

// defer は複数書くことができる。実行される順番は下から上である。
// output is
// Hello
// 111111
// world
// 22222
