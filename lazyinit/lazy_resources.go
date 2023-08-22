package lazyinit

import (
	"fmt"
	"sync"
)

var once sync.Once

var lazyString string

func GetString() string {
	once.Do(func() {
		fmt.Println("Initializing lazyString")

		lazyString = "Hello, World!"
	})
	return lazyString
}
