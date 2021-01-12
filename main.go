package main

import (
	"fmt"

	"github.com/adamzv/go-random/Random"
)

func main() {
	for i := 0; i < 10; i++ {
		// generates numbers from range <0;10>
		fmt.Println(0 + int(Random.RandomGenerator()*11))
	}
}
