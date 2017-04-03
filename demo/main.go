package main

import (
	"fmt"
	"strconv"
)

func FizzBuzz(i int) string {
	if i%3 == 0 {
		return "Fizz"
	}

	if i%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(i)
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(FizzBuzz(i))
	}
}
