/*
* Kayhan Dehghani Mohammadi
* 301243781
* kdehghan@sfu.ca
 */

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"text/scanner"
)

// JSONtoken : to store tokens returned by decodeJSON()
type JSONtoken struct {
	key   uint8
	value string
}

// JSON token definition
const (
	BRACE   = 1
	BRACKET = 2
	COMMA   = 3
	COLON   = 4
	MINUS   = 5
	NUMBER  = 6
	STRING  = 7
)

func main() {
	arguments := os.Args[1:] // ignore the first argument (path to the program)
	if len(arguments) != 1 {
		panic("Error: the program only accepts one argument: the name of the JSON file")
	}

	fileName := arguments[0]
	dataByte, err := ioutil.ReadFile(fileName)

	if err != nil {
		errMsg := "Error: unable to read the file \"" + fileName + "\""
		panic(errMsg)
	}

	dataString := string(dataByte) // convert data from byte stream to string

	// scanJSON(&dataString)
	tokens := scanJSON(&dataString)
	for _, t := range tokens {
		fmt.Println(t.key, ":", t.value)
	}
}

func scanJSON(data *string) []JSONtoken {
	var tokens []JSONtoken
	var token JSONtoken

	var scnr scanner.Scanner
	scnr.Init(strings.NewReader(*data))
	var currentToken rune
	var currentString string
	var stringPattern = regexp.MustCompile(`"(.)*?"`) // regular expression for strings
	var intPattern = regexp.MustCompile("[0-9]+")     // regular expression for numbers

	for currentToken != scanner.EOF {
		currentToken = scnr.Scan() // read tokens
		currentString = scnr.TokenText()

		// check for one character matches
		switch currentString {
		case "{", "}":
			token.key = BRACE
		case "[", "]":
			token.key = BRACKET
		case ",":
			token.key = COMMA
		case ":":
			token.key = COLON
		case "-":
			token.key = MINUS
		}

		// check for pattern matches : string, number
		if stringPattern.MatchString(currentString) {
			token.key = STRING
		} else if intPattern.MatchString(currentString) {
			token.key = NUMBER
		}

		token.value = currentString
		tokens = append(tokens, token)
	}
	return tokens
}
