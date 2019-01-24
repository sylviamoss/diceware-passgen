package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	var password = ""

	for i := 1; i <= 6; i++ {
		index := findDicewareWordIndex()
		word := findDicewareWord(index)
		password = password + word
	}

	fmt.Println("Your six word password: " + password)
}

func findDicewareWordIndex() string {
	var number = ""
	for j := 1; j <= 5; j++ {
		number = number + strconv.Itoa(throwDice())
	}
	return number
}

func throwDice() int {
	return rand.Intn(6)
}

func findDicewareWord(number string) string {
	return "word"
}
