package main

// // type Hangman struct {
// // 	Attempts   uint8
// // 	WordToFind string
// // 	ToRev      []int
// // 	StockChar  []string
// // }

// // isInside returns true if a value (int) is at least one time in a aray (of int) othertwise it returns false
// func isInside(value *int, arr *[]int) bool {
// 	for _, elem := range *arr {
// 		if elem == *value {
// 			return true
// 		}
// 	}
// 	return false
// }

// // isInside returns true if a value (int) is at least one time in a aray (of int) othertwise it returns false
// func isInsideChar(value *string, s *string) bool {
// 	for _, elem := range *s {
// 		if string(elem) == *value {
// 			return true
// 		}
// 	}
// 	return false
// }

// // printWordProgress prints the progess of finding the word
// func printWordProgress(wordToFind *string, toRev *[]int) {
// 	wordToFindLen := len(*wordToFind) - 1

// 	for index, char := range *wordToFind {
// 		if isInside(&index, toRev) {
// 			fmt.Print(strings.ToUpper(string(char)))
// 		} else {
// 			fmt.Print("_")
// 		}

// 		if index != wordToFindLen {
// 			fmt.Print(" ")
// 		}
// 	}
// 	fmt.Println()
// 	fmt.Println()
// }

// // saveGame write to save.txt the status of the game
// func saveGame(status Hangman) {
// 	saveContent := Marshal(status)
// 	err := ioutil.WriteFile("save.txt", saveContent, 0777)
// 	checkError(err)
// 	fmt.Println("Game Saved in save.txt.")
// }

// // Marshal returns the JSON encoding of Hangman
// func Marshal(status Hangman) []byte {
// 	b, err := json.Marshal(status)
// 	checkError(err)
// 	return b
// }

// //Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by Hangman
// func UnMarshal(data []byte) Hangman {
// 	var status Hangman
// 	err := json.Unmarshal(data, &status)
// 	checkError(err)
// 	return status
// }

// // checkError checks if the error is different from nil otherwise displays error
// func checkError(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

// // getSave return the file saves
// func getSave(saveFilename *string) Hangman {
// 	content, err := ioutil.ReadFile(*saveFilename)
// 	checkError(err)
// 	return UnMarshal(content)
// }

// //readFile returns an array of string which is the same as the file (line = line) without useless lines
// func readFile(Filename string) []string {
// 	var source []string
// 	file, err := os.Open(Filename) // opens the .txt
// 	checkError(err)
// 	scanner := bufio.NewScanner(file) // scanner scans the file
// 	scanner.Split(bufio.ScanLines)    // sets-up scanner preference to read the file line-by-line
// 	for scanner.Scan() {              // loop that performs a line-by-line scan on each new iteration
// 		if scanner.Text() != "" {
// 			source = append(source, scanner.Text()) // adds the value of scanner (that contains the characters from StylizedFile) to source
// 		}
// 	}
// 	file.Close() // closes the file
// 	return source
// }

// // Open file to get random word
// func randomWord() string {
// 	words := readFile(flag.Args()[0])
// 	return strings.ToUpper(words[rand.Intn(len(words))])
// }

// func testWord(status *Hangman, UserTry *string) (valid bool, elem string) {
// 	// If the user entry is a char
// 	if len(*UserTry) <= 1 {

// 		AllChar := strings.Join(status.StockChar, "")
// 		// Add the char to StockChar
// 		if !isInsideChar(UserTry, &AllChar) {
// 			status.StockChar = append(status.StockChar, *UserTry)
// 		} else {
// 			fmt.Println("Already try", *UserTry)
// 			return false, "char"
// 		}

// 		if isInsideChar(UserTry, &status.WordToFind) {
// 			for index, char := range status.WordToFind {
// 				if string(char) == *UserTry {
// 					if !isInside(&index, &status.ToRev) {
// 						status.ToRev = append(status.ToRev, index)
// 					}
// 				}
// 			}
// 			return true, "char"
// 		} else {
// 			status.Attempts++
// 			return false, "char"
// 		}
// 	} else { //If it's a word
// 		if *UserTry == status.WordToFind {
// 			return true, "word"
// 		} else {
// 			status.Attempts += 2
// 			return false, "word"
// 		}
// 	}
// }

// func oldmain() {
// 	// set saveFilename as flag string vars
// 	var saveFilename string

// 	// StringVar defines a string flag with specified name, default value, and usage string
// 	// The argument saveFilename points to a string variable in which to store the value of the flag.
// 	flag.StringVar(&saveFilename, "startWith", "", "Specifie the filename of the save to load")

// 	// parse flags from command line
// 	flag.Parse()

// 	fmt.Println(saveFilename)

// 	rand.Seed(time.Now().UnixNano())

// 	var status Hangman

// 	if saveFilename != "" {
// 		status = getSave(&saveFilename)
// 		fmt.Println("Welcome Back, you have", 10-status.Attempts, "attempts remaining.")
// 	} else {

// 		// Check if user provide file containing words
// 		if len(flag.Args()) <= 0 {
// 			fmt.Println("Missing files of words.")
// 			return
// 		}
// 		// Retrieve the chosen random word
// 		wordToFind := randomWord()
// 		for wordToFind == "STOP" {
// 			wordToFind = randomWord()
// 		}

// 		// Number of letter to reveal
// 		reveal := len(wordToFind)/2 - 1

// 		var toRev []int
// 		if reveal > 0 {
// 			var randInt int

// 			for i := 0; i < reveal; i++ {
// 				randInt = rand.Intn(len(wordToFind))

// 				if !isInside(&randInt, &toRev) {
// 					toRev = append(toRev, randInt)
// 				} else {
// 					i--
// 				}
// 			}
// 		}
// 		status = Hangman{Attempts: 0, WordToFind: wordToFind, ToRev: toRev}
// 		fmt.Println("Good luck, you have 10 attempts.")
// 	}
// 	printWordProgress(&status.WordToFind, &status.ToRev)

// 	var UserTry string

// 	for status.Attempts != 10 {
// 		// Get user input
// 		fmt.Print("Choose: ")
// 		fmt.Scanln(&UserTry)
// 		UserTry = strings.ToUpper(UserTry)

// 		if UserTry == "STOP" {
// 			//call save func and exit
// 			saveGame(status)
// 			os.Exit(0)
// 		}

// 		// Test if user input match something in the word
// 		valid, elem := testWord(&status, &UserTry)

// 		if valid && elem == "word" {
// 			fmt.Println("Congrats !")
// 			break
// 		} else if valid && elem == "char" {
// 			printWordProgress(&status.WordToFind, &status.ToRev)

// 			if len(status.ToRev) == len(status.WordToFind) {
// 				fmt.Println("Congrats !")
// 				break
// 			}
// 		} else {
// 			fmt.Println("Not present in the word,", 10-status.Attempts, "attempts remaining")
// 			if status.Attempts > 10 {
// 				status.Attempts = 10
// 			}
// 		}
// 	}

// 	// If the maximum number of try reached
// 	if status.Attempts == 10 {
// 		fmt.Println("Failed")
// 		fmt.Println("Word to find was", status.WordToFind)
// 	}
// }
