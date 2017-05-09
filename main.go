package main

import (
	"fmt"
)

func doSome(i int) {
	fmt.Printf("%d, hello world test2\n", i)

}

func main() {
	var i int
	for i = 0; i < 10; i++ {
		doSome(i)
	}
}
