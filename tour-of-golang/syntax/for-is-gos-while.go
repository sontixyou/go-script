package main

import "fmt"

// whileのように書きたければ、セミコロンを省略することで書ける
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	// これで無限ループになる。おもろい
	// for {
	// 	sum += 1
	// 	fmt.Println(sum)
	// }
	fmt.Println(sum)
}
