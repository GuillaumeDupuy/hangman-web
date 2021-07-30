package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Hangman struct {
	Attempts   uint8
	WordToFind string
	ToRev      []int
}

type HangmanData struct {
	Attempts       uint8
	WordToFind     string
	WordInProgress string
}

func pageChecker(w http.ResponseWriter, r *http.Request, page string) bool {
	if r.URL.Path != page { // Checks if we are in  /ascii-art
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

	// Parsing data
	// file.Execute(w, data)

	// Render template
	if error := file.Execute(w, file); error != nil {
		log.Fatal(error)
	}
}

// /hangman page
func hangmanHandler(w http.ResponseWriter, r *http.Request) {
	if pageChecker(w, r, "/hangman") {
		return
	}

	file, error := template.ParseFiles("./templates/hangman.html")

	if error != nil {
		log.Fatal(error)
	}

	// Render template
	if error := file.Execute(w, file); error != nil {
		log.Fatal(error)
	}
}

func main() {
	// Serving templates files
	filesServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", filesServer))

	http.HandleFunc("/", rootHandler)           // handle root (main) webpage
	http.HandleFunc("/hangman", hangmanHandler) // handle ascii-art page

	fmt.Println("Starting server at port 3000")
	fmt.Println("Go on http://localhost:3000") // Prints the link of the website on the command prompt
	fmt.Println("\nTo shutdown the server and exit the code hit \"crtl+C\"")
	if err := http.ListenAndServe(":3000", nil); err != nil { // Launches the server on port 8080 if port 8080 is not already busy, else quit
		log.Fatal(err)
	}
}
