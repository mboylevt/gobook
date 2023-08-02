package main

import "fmt"

const value = 1

func main() {
	message := "Hi � and �"
	rMs := []rune(message)
	fmt.Println(string(rMs[3]))
}
