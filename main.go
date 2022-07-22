package main

import (
	"fmt"
)

func time() int64 {
	var i, x int64 = 0, 0
	for i = 0; i < 1000000000; i++ {
		x = x + i
	}
	return x
}

func main() {
	fmt.Println(time())
}
