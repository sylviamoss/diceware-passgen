package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	var password = ""

	for i := 1; i <= 6; i++ {
		index := findDicewareWordIndex()
		fmt.Println(index)
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
	var number = 0

	for number == 0 {
		number = rand.Intn(6)
	}

	return number
}

func findDicewareWord(number string) string {
	fmt.Println(number)
	file, err := os.Open("diceware_words/" + number + ".txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

// func createWordsFiles() {
// 	file, err := os.Open("diceware_words.txt")

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {
// 		words := strings.Fields(scanner.Text())
// 		f, _ := os.Create("diceware_words/" + words[0] + ".txt")
// 		f.WriteString(words[1])
// 		f.Sync()
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// }
