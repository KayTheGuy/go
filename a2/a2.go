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
	BRACE     = 1
	BRACKET   = 2
	COMMA     = 3
	COLON     = 4
	MINUS     = 5
	NUMBER    = 6
	NULTRUFLS = 7 // null, true, or false
	STRING    = 8
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

	tokens := scanJSON(&dataString) // get JSON tokens
	// for _, t := range tokens {
	// 	fmt.Println(t.key, ":", t.value)
	// }
	formatHTML(&tokens)
}

func scanJSON(data *string) []JSONtoken {
	var tokens []JSONtoken
	var token JSONtoken

	var scnr scanner.Scanner
	scnr.Init(strings.NewReader(*data))
	var currentToken rune
	var currentString string
	// regular expression for strings
	var stringPattern = regexp.MustCompile(`"(.)*?"`)
	// regular expression for numbers
	// NOTE: existence of even one digit suffice to recognize the token as number
	// 		 because it is evaluated after string and the problem assumes the input JSON is valid
	var numPattern = regexp.MustCompile("[0-9]+")

	currentToken = scnr.Scan() // read first token
	for currentToken != scanner.EOF {
		currentString = scnr.TokenText()
		currentToken = scnr.Scan() // read the remaining tokens including EOF

		// find type for JSONtoken type
		switch {
		case currentString == "{", currentString == "}":
			token.key = BRACE
		case currentString == "[", currentString == "]":
			token.key = BRACKET
		case currentString == ",":
			token.key = COMMA
		case currentString == ":":
			token.key = COLON
		case currentString == "-":
			token.key = MINUS
		case currentString == "null", currentString == "true", currentString == "false":
			token.key = NULTRUFLS
		case stringPattern.MatchString(currentString):
			token.key = STRING
		case numPattern.MatchString(currentString):
			token.key = NUMBER
		}
		token.value = currentString
		tokens = append(tokens, token)
	}
	return tokens
}

func formatHTML(tokens *[]JSONtoken) {
	// print the HTML headers
	fmt.Println("<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 4.01//EN\" \"http://www.w3.org/TR/html4/strict.dtd\">")
	fmt.Println("<html>\n<head>\n<meta http-equiv=\"Content-Type\" content=\"text/html; charset=UTF-8\"/>")
	fmt.Println("<title>JSON to HTML Formater</title>\n<link rel=\"stylesheet\" href=\"style.css\"/>\n</head><body>")
	fmt.Println("<span style=\"font-family:monospace; white-space:pre\">")

	// begining of JSON code
	// TODO: finalize indentation and colors
	for _, token := range *tokens {
		switch token.key {
		case BRACE:
			fmt.Printf("<br><span style=\"color:red\">%v</span>&nbsp;", token.value)
		case BRACKET:
			fmt.Printf("<span style=\"color:#ffa500\">%v</span>", token.value)
		case COMMA:
			fmt.Printf("<span style=\"color:#2cad6b\">&nbsp;%v <br>&nbsp;</span>", token.value)
		case COLON:
			fmt.Printf("<span style=\"color:#00ffff\">&nbsp;&nbsp;&nbsp; %v &nbsp;&nbsp;&nbsp;</span>", token.value)
		case MINUS:
			fmt.Printf("<span style=\"color:#ff0000\">%v</span>", token.value)
		case NUMBER:
			fmt.Printf("<span style=\"color:#477ec7\">%v</span>", token.value)
		case NULTRUFLS:
			fmt.Printf("<span style=\"color:#b22222\">%v</span>", token.value)
		case STRING:
			fmt.Printf("<span style=\"color:black\">%v", token.value)
			// TODO: change color for scape characters
			// TODO: replace html espcial characters
		}
	}

	// end of JSOC code

	fmt.Println("</span></body></html>")

}
