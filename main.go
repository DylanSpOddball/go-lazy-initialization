package main

import (
	"fmt"

	"github.com/DylanSpOddball/go-lazy-initialization/lazyinit"
)

// harness for demonstrating sync.Once usage taken from example in godocs
// https://pkg.go.dev/sync#example-Once
func main() {
	numIterations := 100

	done := make(chan bool)
	for i := 0; i < numIterations; i++ {
		// calling lazyinit.GetString() will only initialize the string once, even when called from multiple goroutines
		go func(iteration int) {
			str := lazyinit.GetString()
			fmt.Println(str, iteration)
			done <- true
		}(i)
	}

	for i := 0; i < numIterations; i++ {
		<-done
	}
}
