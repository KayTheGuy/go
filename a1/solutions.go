/*
* Kayhan Dehghani Mohammadi
* 301243781
* kdehghan@sfu.ca
 */

package a1

import (
	"errors"
	"io/ioutil"
	"math"
	"reflect"
	"strconv"
	"strings"
)

//=======================================================================
// Problem 1: find the number of primes less than, or equal to <input in>
//=======================================================================

func isPrime(num int) bool {
	// if num is even and not equal to 2 return false
	if num%2 == 0 && num != 2 {
		return false
	}
	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func countPrimes(n int) int {
	count := 0
	if n < 2 {
		return count
	}
	for i := 2; i <= n; i++ {
		// if the number is prime increment the count
		if isPrime(i) {
			count++
		}
	}
	return count
}

//=======================================================================
// Problem 2: read a text file <filename>
// return a <map[string]int> contaning the counts of each word
// returns nil if unable to read the file
//=======================================================================

func countStrings(filename string) map[string]int {
	// read the file
	data, ok := ioutil.ReadFile(filename)
	wordsCount := make(map[string]int) // words and their corresponding counts in the text

	if ok != nil {
		return nil
	}
	text := string(data)          // convert data from byte to string
	words := strings.Fields(text) // extract words between whitespaces

	for _, word := range words {
		_, ok := wordsCount[word]
		if ok {
			// if the word exists in the map increment its count
			wordsCount[word]++
		} else {
			// initialize its count if new word
			wordsCount[word] = 1
		}
	}
	return wordsCount
}

//=======================================================================
// Problem 3: Different Functions and Methods for type Time24
//=======================================================================

// Time24 : 0 <= hour < 24, 0 <= minute < 60, 0 <= second < 60
type Time24 struct {
	hour, minute, second uint8
}

// returns true if <a> and <b> are exactly the same time
func equalsTime24(a, b Time24) bool {
	isEqual := false
	if a.hour == b.hour && a.minute == b.minute && a.second == b.second {
		isEqual = true
	}
	return isEqual
}

// returns true if <a> is strictly less than <b>
func lessThanTime24(a, b Time24) bool {
	isLess := false
	if (a.hour < b.hour) || (a.hour == b.hour && a.minute < b.minute) || (a.hour == b.hour && a.minute == b.minute && a.second < b.second) {
		isLess = true
	}
	return isLess
}

// converts <t> to human-readable string
func (t Time24) String() string {
	h := formatTime24Element(strconv.Itoa(int(t.hour)))
	m := formatTime24Element(strconv.Itoa(int(t.minute)))
	s := formatTime24Element(strconv.Itoa(int(t.second)))
	return h + ":" + m + ":" + s
}

// helper for t.String() method
// if <t> is between 0 and 10 returns "0t"
func formatTime24Element(t string) string {
	if len(t) == 1 {
		t = "0" + t
	}
	return t
}

// returns true if <t> is a valid Time24 objec
// <t> is valid if 0 <= t.hour < 24, 0 <= t.minute < 60, 0 <= t.second < 60
func (t Time24) validTime24() bool {
	isValid := false
	if t.hour >= 0 && t.hour < 24 && t.minute >= 0 && t.minute < 60 && t.second >= 0 && t.second < 60 {
		isValid = true
	}
	return isValid
}

// returns the smallest Time24 in the slice <times>
// returns an error and Time24{0, 0, 0} if slice is empty
func minTime24(times []Time24) (Time24, error) {
	sliceLength := len(times) // lenght of slice
	if sliceLength == 0 {
		return Time24{0, 0, 0}, errors.New("Error: the argument should not be empty slice")
	}
	// find the minTime24
	minTime := times[0] // initialize minTime24
	for i := 1; i < sliceLength; i++ {
		if lessThanTime24(times[i], minTime) {
			minTime = times[i] // update the minTime24
		}
	}
	return minTime, nil
}

//==========================================================================================
// Problem 4: use linear search to return the first index location of <x> in the slice <lst>
// return -1 if element not found
// works for int and string
// return -1 if type of <x> and <lst> is the same but not supported (i.e. not int or string)
// Panic if the type of <x> is not the same as the type of the elements in <slc>
//==========================================================================================

func linearSearch(x interface{}, lst interface{}) int {
	// get the type of value <x> and slice <lst> elements
	xType := reflect.TypeOf(x)
	lstType := reflect.TypeOf(lst).Elem()

	// panic if types of <x> and <lst> differ
	if xType != lstType {
		errorMessage := "Error: Type mismatch in arguments of linearSearch()\n First argument is of type " + strings.ToUpper(xType.String()) + " and slice elements is second arguments are of type " + strings.ToUpper(lstType.String())
		panic(errorMessage)
	}

	switch xType.Kind() {
	case reflect.Int:
		intX, intXOk := x.(int)
		intLst, intLstOk := lst.([]int)
		if intXOk && intLstOk {
			return intLinearSearch(intX, &intLst)
		}
	case reflect.String:
		stringX, stringXOk := x.(string)
		stringLst, stringLtOk := lst.([]string)
		if stringXOk && stringLtOk {
			return stringLinearSearch(stringX, &stringLst)
		}
	}
	return -1
}

// helper function for linear search for int values
func intLinearSearch(x int, lst *[]int) int {
	for idx, val := range *lst {
		if val == x {
			return idx // found the value in the slice
		}
	}
	return -1
}

// helper function for linear search for string values
func stringLinearSearch(x string, lst *[]string) int {
	for idx, val := range *lst {
		if val == x {
			return idx // found the value in the slice
		}
	}
	return -1
}

//=====================================================
// Problem 5: return all the bit sequences of length n
//		      return empty slice if <n> 0 or negative
// Example:
// 		allBitSeqs(1) returns [[0] [1]].
// 		allBitSeqs(2) returns [[0 0] [0 1] [1 0] [1 1]]
//=====================================================
func allBitSeqs(n int) [][]int {
	// return empty slice upon invalid slice size
	if n <= 0 {
		return [][]int{}
	}
	// define the length of bit sequences set
	// <length> = 2 to the power of n (all permutations of sequences of size n)
	length := int(math.Pow(2, float64(n)))

	// make a slice of lenght = <lenght> and capacity = <length>
	var allBits = make([][]int, length)

	// allocate a 1d slice of size <n> to each row of <allBits>
	for row := range allBits {
		allBits[row] = make([]int, n)
	}

	base2Value := 0
	// make all permutations of 0 and 1
	for row := length - 1; row >= 0; row-- {
		// value of <row> represents the value of sequence in base 2
		// example: <row> = 5 represents the sequence 101
		base2Value = row
		for column := n - 1; column >= 0; column-- {
			allBits[row][column] = base2Value % 2 // assign 1 or 0 to corresponding bit position
			base2Value = base2Value / 2           // prepare for the next postion
		}
	}
	return allBits
}
