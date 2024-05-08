package main

import "fmt"

// func returns another func
func generator() func() int {
	var i = 0
	return func() int {
		i++
		return i
	}
}

func main() {
	numGen := generator()
	for i := 0; i < 10; i++ {
		fmt.Print(numGen(), "\t")
	}
}
