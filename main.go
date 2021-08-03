package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type HangmanData struct {
	Attempts       uint8
	WordToFind     string
	WordInProgress []string
	ToRev          []int
	Status         string
}

var data HangmanData

// checkError checks if the error is different from nil otherwise displays error
func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

// isInside returns true if a value (int) is at least one time in a aray (of int) othertwise it returns false
func isInsideChar(value string, s string) bool {
	for _, elem := range s {
		if string(elem) == value {
			return true
		}
	}
	return false
}

// isInside returns true if a value (int) is at least one time in a aray (of int) othertwise it returns false
func isInside(value int, arr []int) bool {
	for _, elem := range arr {
		if elem == value {
			return true
		}
	}
	return false
}

//readFile returns an array of string which is the same as the file (line = line) without useless lines
func readFile(Filename string) []string {
	var source []string
	file, err := os.Open(Filename) // opens the .txt
	checkError(err)
	scanner := bufio.NewScanner(file) // scanner scans the file
	scanner.Split(bufio.ScanLines)    // sets-up scanner preference to read the file line-by-line
	for scanner.Scan() {              // loop that performs a line-by-line scan on each new iteration
		if scanner.Text() != "" {
			source = append(source, scanner.Text()) // adds the value of scanner (that contains the characters from StylizedFile) to source
		}
	}
	file.Close() // closes the file
	return source
}

func pageChecker(w http.ResponseWriter, r *http.Request, page string) bool {
	if r.URL.Path != page { // Checks if we are in the right page
		http.Error(w, "404 not found.", http.StatusNotFound) // Sends a 404 error (page not found)
		return true
	}
	return false
}

// / page
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if pageChecker(w, r, "/") {
		return
	}

	file, error := template.ParseFiles("./templates/index.html")

	if error != nil {
		log.Fatal(error)
	}

	// data := HangmanData{Attempts: 10, WordToFind: "HELLO", WordInProgress: "H_LL_"}

	toRevToWord(data.WordInProgress, data.WordToFind, data.ToRev)

	if strings.Join(data.WordInProgress, "") == data.WordToFind {
		fmt.Println("Finito")
	}

	if data.Attempts == 0 || data.Attempts > 10 {
		fmt.Println("Failed")
		fmt.Println("The word was", data.WordToFind)
		data.Attempts = 0
	}

	// Parsing data
	file.Execute(w, data)

	// Render template
	// if error := file.Execute(w, file); error != nil {
	// 	log.Fatal(error)
	// }
}

func readForm(w http.ResponseWriter, r *http.Request) string {
	r.ParseForm()

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
	}

	return strings.ToUpper(r.FormValue("letter"))
}

// /hangman page
func hangmanHandler(w http.ResponseWriter, r *http.Request) {
	if pageChecker(w, r, "/hangman") {
		return
	}

	userInput := readForm(w, r)

	if r.Method == "POST" {
		if len(userInput) == 1 {
			if isInsideChar(userInput, data.WordToFind) {
				for index, char := range data.WordToFind {
					if string(char) == userInput {
						if !isInside(index, data.ToRev) {
							data.ToRev = append(data.ToRev, index)
						}
					}
				}
			} else {
				data.Attempts--
			}
		} else {
			// User input is not a letter
			data.Attempts--
		}

		// Rediect to the main page so the /hangman page is only here to call the hangman with a new letter
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

// toRevToWord replace '_' char with the right char
func toRevToWord(wordInProgress []string, wordToFind string, toRev []int) {
	for _, index := range toRev {
		wordInProgress[index] = string([]rune(wordToFind)[index])
	}
}

func getRandomWord() string {
	words := readFile(os.Args[1])
	return strings.ToUpper(words[rand.Intn(len(words))])
}

func buildWordInProgress(wordToFind string) []string {
	wordInProgress := make([]string, len(wordToFind))

	for index := range wordInProgress {
		wordInProgress[index] = "_"
	}

	return wordInProgress
}

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Println("Missing word file")
		return
	}

	rand.Seed(time.Now().UnixNano())

	// Open file to get word
	wordToFind := getRandomWord()

	// Number of letter to reveal
	reveal := len(wordToFind)/2 - 1

	var toRev []int
	if reveal > 0 {
		var randInt int

		for i := 0; i < reveal; i++ {
			randInt = rand.Intn(len(wordToFind))

			if !isInside(randInt, toRev) {
				toRev = append(toRev, randInt)
			} else {
				i--
			}
		}
	}

	// Build word display
	wordInProgress := buildWordInProgress(wordToFind)
	toRevToWord(wordInProgress, wordToFind, toRev)

	// store values
	data = HangmanData{Attempts: 10, WordToFind: wordToFind, WordInProgress: wordInProgress, ToRev: toRev}

	// Serving static files
	filesServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", filesServer))

	http.HandleFunc("/", rootHandler)           // handle root (main) webpage
	http.HandleFunc("/hangman", hangmanHandler) // handle hangman logic

	fmt.Println("Starting server at port 3000")
	fmt.Println("Go on http://localhost:3000") // Prints the link of the website on the command prompt
	fmt.Println("\nTo shutdown the server and exit the code hit \"crtl+C\"")
	if err := http.ListenAndServe(":3000", nil); err != nil { // Launches the server on port 3000 if port 3000 is not already busy, else quit
		log.Fatal(err)
	}
}
