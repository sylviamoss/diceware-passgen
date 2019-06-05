package main

import (
	"bufio"
	"crypto/rand"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type PasswordResponse struct {
	Words    string `json:"words"`
	Password string `json:"password"`
}

func main() {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/{lang}", GenerateDicewarePassword).Methods("GET")

	buildHandler := http.FileServer(http.Dir("frontend/build/"))
	router.PathPrefix("/").Handler(buildHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), handlers.LoggingHandler(os.Stdout, router))
}

func GenerateDicewarePassword(w http.ResponseWriter, r *http.Request) {
	var lang = mux.Vars(r)["lang"]
	var words = ""
	var password = ""
	var response PasswordResponse

	for i := 1; i <= 6; i++ {
		index := findDicewareWordIndex()
		word := findDicewareWord(index, lang)
		words = words + word + " "
		password = password + word
	}

	response.Words = strings.TrimSpace(words)
	response.Password = password

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func findDicewareWordIndex() string {
	var number = ""
	for j := 1; j <= 5; j++ {
		number = number + strconv.FormatInt(throwDice(), 10)
	}
	return number
}

func throwDice() int64 {
	var number int64 = 0

	for number == 0 {
		nBig, err := rand.Int(rand.Reader, big.NewInt(7))
		if err != nil {
			panic(err)
		}
		number = nBig.Int64()
	}

	return number
}

func findDicewareWord(number string, lang string) string {
	file, err := os.Open("diceware_words_" + lang + "/" + number + ".txt")

	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()
		t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
		transformedWord, _, _ := transform.String(t, word)
		return transformedWord
	}
	return ""
}

func createWordsFiles() {
	file, err := os.Open("diceware_words_pt.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		f, _ := os.Create("diceware_words_pt/" + words[0] + ".txt")
		f.WriteString(words[1])
		f.Sync()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
