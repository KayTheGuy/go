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

// JSONtoken : to store tokens returned by scanJSON()
type JSONtoken struct {
	key    uint8
	lexeme string
}

// constant values for JSONtoken.key
const (
	BRACE     uint8 = 1
	BRACKET   uint8 = 2
	COMMA     uint8 = 3
	COLON     uint8 = 4
	MINUS     uint8 = 5
	NUMBER    uint8 = 6
	NULTRUFLS uint8 = 7 // null, true, or false
	STRING    uint8 = 8
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

	dataString := string(dataByte)  // convert data from byte stream to string
	tokens := scanJSON(&dataString) // get JSON tokens
	formatHTML(&tokens)             // format in HTML
}

func scanJSON(data *string) []JSONtoken {
	var tokens []JSONtoken
	var token JSONtoken
	var currentToken rune
	var currentString string

	var scnr scanner.Scanner
	scnr.Init(strings.NewReader(*data))

	// regular expression for strings
	stringPattern := regexp.MustCompile(`"(.)*?"`)

	/* regular expression for numbers
	NOTE: existence of even one digit suffice to recognize the token as number
	because it is evaluated after string and the problem assumes the input JSON is valid*/
	numPattern := regexp.MustCompile("[0-9]+")

	currentToken = scnr.Scan() // read first token
	for currentToken != scanner.EOF {
		currentString = scnr.TokenText()

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
		token.lexeme = currentString
		tokens = append(tokens, token)

		currentToken = scnr.Scan() // read the remaining tokens including EOF
	}
	return tokens
}

func formatHTML(tokens *[]JSONtoken) {
	/* replacer for HTML special characters:
	   < with &lt;  > with &gt;  & with &amp;
	   " with &quot; ' with &apos; Space with &nbsp; */
	htmlReplacer := strings.NewReplacer(
		"\"", "&quot;",
		"'", "&apos;",
		"&", "&amp;",
		">", "&gt;",
		"<", "&lt;",
		" ", "&nbsp;",
	)
	/* replacer to change color of JSON escape characters within string:
	e.g. : \n with <span style=\"color:#FF8C00\">\n</span>
	*/
	escapeReplacer := strings.NewReplacer(
		"\\n", "<span style=\"color:#FF8C00\">\\n</span>",
		"\\b", "<span style=\"color:#FF8C00\">\\b</span>",
		"\\f", "<span style=\"color:#FF8C00\">\\f</span>",
		"\\r", "<span style=\"color:#FF8C00\">\\r</span>",
		"\\t", "<span style=\"color:#FF8C00\">\\t</span>",
		"\\u", "<span style=\"color:#FF8C00\">\\u</span>",
		"\\/", "<span style=\"color:#FF8C00\">\\/</span>",
		"\\\\", "<span style=\"color:#FF8C00\">\\\\</span>",
		"\\\"", "<span style=\"color:#FF8C00\">\\\"</span>",
	)

	// value used for indenting nested brackets and commas
	numOfIndent := -1
	// print the HTML headers
	fmt.Println("<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 4.01//EN\" \"http://www.w3.org/TR/html4/strict.dtd\">")
	fmt.Println("<html xmlns=\"http://www.w3.org/1999/xhtml\">\n<head><meta http-equiv=\"Content-Type\" content=\"text/html;charset=ISO-8859-8\"/>")
	fmt.Println("<title>JSON to HTML Formater</title>\n</head><body style=\"background-color:#DEDEDC\">")
	fmt.Println("<p><span style=\"font-family:monospace; white-space:pre\">")

	// start of JSON code
	for _, token := range *tokens {
		switch token.key {
		case BRACE:
			fmt.Printf("<br/>")
			if token.lexeme == "{" {
				numOfIndent++
			}
			for i := 0; i < numOfIndent; i++ {
				fmt.Printf("&nbsp;&nbsp;")
			}
			fmt.Printf("<span style=\"color:red\">%v</span>", token.lexeme)
			if token.lexeme == "}" {
				numOfIndent--
			}
		case BRACKET:
			fmt.Printf("<span style=\"color:#FFFF00\">%v</span>", token.lexeme)
		case COMMA:
			fmt.Printf("<span style=\"color:#FF00FF\">%v</span><br/>", token.lexeme)
			for i := 0; i < numOfIndent; i++ {
				fmt.Printf("&nbsp;&nbsp;&nbsp;")
			}
		case COLON:
			fmt.Printf("<span style=\"color:#3BB9FF\">%v &nbsp;</span>", token.lexeme)
		case MINUS:
			fmt.Printf("<span style=\"color:#F88017\">%v</span>", token.lexeme)
		case NUMBER:
			fmt.Printf("<span style=\"color:#2B65EC\">%v</span>", token.lexeme)
		case NULTRUFLS:
			fmt.Printf("<span style=\"color:#b22222\">%v</span>", token.lexeme)
		case STRING:
			// replace html espcial characters and change color for escape chars
			fmt.Printf("<span style=\"color:black\">%v</span>", escapeReplacer.Replace(htmlReplacer.Replace(token.lexeme)))
		}
	}
	// end of JSOC code

	fmt.Println("</span></p></body></html>") // close HTML header tags
}
