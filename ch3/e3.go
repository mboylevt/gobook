package main

import "fmt"

const value = 1

func main() {
	var b byte = 127
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615
	fmt.Println(b + 1)
	fmt.Println(smallI + 1)
	fmt.Println(bigI + 1)
}
