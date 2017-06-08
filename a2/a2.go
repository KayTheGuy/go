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
	// temporary string holder
	tmp := ""
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
	e.g. : \n with <span style=\"color:#cefc53\">\n</span>
	*/
	escapeReplacer := strings.NewReplacer(
		"\\n", "<span style=\"color:#cefc53\">\\n</span>",
		"\\b", "<span style=\"color:#cefc53\">\\b</span>",
		"\\f", "<span style=\"color:#cefc53\">\\f</span>",
		"\\r", "<span style=\"color:#cefc53\">\\r</span>",
		"\\t", "<span style=\"color:#cefc53\">\\t</span>",
		"\\/", "<span style=\"color:#cefc53\">\\/</span>",
		"\\\\", "<span style=\"color:#cefc53\">\\\\</span>",
	)

	// value used for indenting nested brackets and commas
	numOfIndent := -1
	// print the HTML headers
	fmt.Println("<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 4.01//EN\" \"http://www.w3.org/TR/html4/strict.dtd\">")
	fmt.Println("<html xmlns=\"http://www.w3.org/1999/xhtml\">")
	fmt.Println("<head><meta http-equiv=\"Content-Type\" content=\"text/html;charset=ISO-8859-8\"/>")
	fmt.Println("<title>JSON Colorizer</title>\n</head><body style=\"background-color:#8e8e8e\">")
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
			fmt.Printf("<span style=\"color:#ffffff\">%v</span>", token.lexeme)
			if token.lexeme == "}" {
				numOfIndent--
			}
		case BRACKET:
			fmt.Printf("<br/>")
			if token.lexeme == "[" {
				numOfIndent++
			}
			for i := 0; i < numOfIndent; i++ {
				fmt.Printf("&nbsp;&nbsp;")
			}
			fmt.Printf("<span style=\"color:#fc5366\">%v</span>", token.lexeme)
			if token.lexeme == "]" {
				numOfIndent--
			}
		case COMMA:
			fmt.Printf("<span style=\"color:#fceb53\">%v</span><br/>", token.lexeme)
			for i := 0; i < numOfIndent; i++ {
				fmt.Printf("&nbsp;&nbsp;&nbsp;")
			}
		case COLON:
			fmt.Printf("<span style=\"color:#53fcd1\">%v &nbsp;</span>", token.lexeme)
		case MINUS:
			fmt.Printf("<span style=\"color:#F88017\">%v</span>", token.lexeme)
		case NUMBER:
			fmt.Printf("<span style=\"color:#2B65EC\">%v</span>", token.lexeme)
		case NULTRUFLS:
			fmt.Printf("<span style=\"color:#931747\">%v</span>", token.lexeme)
		case STRING:
			// replace html espcial characters and change color for escape chars
			tmp = escapeReplacer.Replace(htmlReplacer.Replace(token.lexeme))
			// colorize hex digits and replace double quotes
			tmp = hexColorize(&tmp)
			// colorize escaped quotation marks
			tmp = escapedQuoteColorize(&tmp)
			// print result string
			fmt.Printf("<span style=\"color:black\">%v</span>", tmp)
		}
	}
	// end of JSOC code

	fmt.Println("</span></p></body></html>") // close HTML header tags
}

// helper function for formatHTML
// colorizes the hex decimal numbers within the argument
// assumes string has a valid hex number of digits : 4 digits after \u character
func hexColorize(s *string) string {
	hexStr := ""       // temporary string holding the hex number
	replacingStr := "" // temporary string holding the colorized hex number
	remainingStr := "" // holds the reamining substring
	newStr := *s
	idx := 0 // points to the starting of hex number in each remaining substring
	remainingIdx := 0
	length := len(*s)
	if strings.Contains(*s, "\\u") {
		newStr = ""
		for remainingIdx <= length {
			// remaining of the argument
			remainingStr = (*s)[remainingIdx:length]
			// index of the first occurrence of '\u' in the ramining substring
			idx = strings.Index(remainingStr, "\\u")
			if idx >= 0 { // if s contains hex number
				hexStr = remainingStr[idx : idx+6] // substring holding the hex number
				replacingStr = "<span style=\"color:#fcbb53\">" + hexStr + "</span>"
				newStr = newStr + remainingStr[0:idx] + replacingStr
				// read from the end of the last hex number (each hex number is 6 characters long)
				remainingIdx = remainingIdx + idx + 6
			} else {
				// add the rest of string after the last hex number
				newStr += (*s)[remainingIdx:length]
				break
			}
		}
	}

	return newStr
}

// helper function for formatHTML
// colorizes the '\*' ('\&quot;') in the argument
func escapedQuoteColorize(s *string) string {
	subStr := ""
	replacingStr := ""
	newStr := *s
	idx := strings.Index(*s, "\\&quot;")
	if idx >= 0 { // if s contains hex number
		subStr = (*s)[idx : idx+7] // substring holding the hex number
		replacingStr = "<span style=\"color:#cefc53\">" + subStr + "</span>"
		newStr = strings.Replace(*s, subStr, replacingStr, -1) // replace all occurrences
	}
	return newStr
}
