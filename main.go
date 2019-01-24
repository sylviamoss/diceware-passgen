package main

import (
	"bufio"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GenerateDicewarePassword).Methods("GET")
	http.ListenAndServe(":"+os.Getenv("PORT"), handlers.LoggingHandler(os.Stdout, router))
}

func GenerateDicewarePassword(w http.ResponseWriter, r *http.Request) {
	var password = ""

	for i := 1; i <= 6; i++ {
		index := findDicewareWordIndex()
		word := findDicewareWord(index)
		password = password + word + " "
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(password)
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

	randSource := rand.NewSource(time.Now().UnixNano())
	randomGen := rand.New(randSource)

	for number == 0 {
		number = randomGen.Intn(6)
	}

	return number
}

func findDicewareWord(number string) string {
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
