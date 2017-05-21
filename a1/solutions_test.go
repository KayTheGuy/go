/*
* Kayhan Dehghani Mohammadi
* 301243781
* kdehghan@sfu.ca
 */

package a1

import (
	"reflect"
	"testing"
)

// Input and Output pair for testing <countPrimes>
type INTPair struct {
	input, output int
}

//=======================================================================
// Problem 1: find the number of primes less than, or equal to <input in>
//=======================================================================
func TestCountPrimes(t *testing.T) {
	testVals := [8]INTPair{
		{10000, 1229},
		{5, 3},
		{-6, 0},
		{0, 0},
		{1, 0},
		{13, 6},
		{19, 8},
		{29, 10},
	}

	for _, val := range testVals {
		if val.output != countPrimes(val.input) {
			t.Fatalf("Error: countPrimes: Expected %d as output but received %d", val.output, countPrimes(val.input))
		}
	}
}

//=======================================================================
// Problem 2: read a text file <filename>
// return a <map[string]int> contaning the counts of each word
//=======================================================================

func TestCountStrings(t *testing.T) {
	// test the erroneous case
	wrongMap := countStrings("unvalidFileName.txt")
	if wrongMap != nil {
		t.Fatalf("Error: countStrings: Expected nil response but the response was: %v", wrongMap)
	}

	// test the <p2_input.txt>
	retunedMap := countStrings("p2_input1.txt")
	expectedMap := map[string]int{"The": 1, "the": 1, "big": 3, "dog": 1, "ate": 1, "apple": 1}

	// compare the expected and returned maps
	equal := reflect.DeepEqual(retunedMap, expectedMap)
	if !equal {
		t.Fatalf("Error: countStrings: The expected map and the returned map by 'countStrings()' differ in at least one element\nTested file was: p2_input1.txt")
	}

	// test the <p2_input2.txt>
	retunedMap = countStrings("p2_input2.txt")
	expectedMap2 := map[string]int{"cmpt": 4, "kayhan": 1, "test": 6, "383": 3}

	// compare the expected and returned maps
	equal2 := reflect.DeepEqual(retunedMap, expectedMap2)
	if !equal2 {
		t.Fatalf("Error: countStrings: The expected map and the returned map by 'countStrings()' differ in at least one element\nTested file was: p2_input2.txt")
	}

	// test empty file
	retunedMap = countStrings("p2_input3.txt")

	if len(retunedMap) != 0 {
		t.Fatalf("Error: countStrings: The expected map and the returned map by 'countStrings()' differ in at least one element\nTested file was: p2_input3.txt")
	}

}

//=======================================================================
// Problem 3: Different Functions and Methods for type Time24
//=======================================================================

func TestEqualsTime24(t *testing.T) {
	t1 := Time24{0, 1, 05}
	t2 := Time24{00, 01, 5}
	t3 := Time24{000, 4, 5}
	t4 := Time24{35, 65, 34}

	if !equalsTime24(t1, t1) {
		t.Fatalf("Error: equalsTime24:  expected t1 be equal to itself: t1 was: %v ", t1.String())
	}
	if !equalsTime24(t1, t2) {
		t.Fatalf("Error: equalsTime24:  expected t1 be equal to t2: t1 was: %v and t2 was: %v", t1.String(), t2.String())
	}
	if equalsTime24(t1, t3) {
		t.Fatalf("Error: equalsTime24:  expected t1 and t3 be different: t1 was: %v and t3 was: %v", t1.String(), t3.String())
	}
	if equalsTime24(t1, t4) {
		t.Fatalf("Error: equalsTime24:  expected t1 and t4 be different: t1 was: %v and t3 was: %v", t1.String(), t4.String())
	}
}

func TestLessThanTime24(t *testing.T) {
	t1 := Time24{0, 1, 05}
	t2 := Time24{00, 01, 5}
	t3 := Time24{000, 4, 5}
	t4 := Time24{35, 65, 34}

	if lessThanTime24(t1, t1) {
		t.Fatalf("Error: lessThanTime24: expected t1 not be less than itself: t1 was: %v ", t1.String())
	}
	if lessThanTime24(t1, t2) {
		t.Fatalf("Error: lessThanTime24: expected t1 not be less than t2: t1 was: %v and t2 was: %v", t1.String(), t2.String())
	}
	if !lessThanTime24(t1, t3) {
		t.Fatalf("Error: lessThanTime24: expected t1 be less than t3: t1 was: %v and t3 was: %v", t1.String(), t3.String())
	}
	if !lessThanTime24(t2, t4) {
		t.Fatalf("Error: lessThanTime24: expected t2 be less than t4: t2 was: %v and t4 was: %v", t2.String(), t4.String())
	}
}

func TestString(t *testing.T) {
	t1 := Time24{0, 1, 05}
	t2 := Time24{00, 01, 5}
	t3 := Time24{000, 44, 5}

	s1 := t1.String()
	s2 := t2.String()
	s3 := t3.String()

	if s1 != "00:01:05" || s2 != "00:01:05" || s3 != "00:44:05" {
		t.Fatalf("Error: Sting() method is not working as expected")
	}
}

func TestValidTime24(t *testing.T) {
	t1 := Time24{0, 1, 05}
	t2 := Time24{60, 01, 5}
	t3 := Time24{0000, 44, 5}
	t4 := Time24{0000, 200, 80}

	if !t1.validTime24() {
		t.Fatalf("Error: validTime24() method : expected t1 be valid. t1 was: %v", t1.String())
	}
	if t2.validTime24() {
		t.Fatalf("Error: validTime24() method : expected t2 be invalid. t2 was: %v", t2.String())
	}
	if !t3.validTime24() {
		t.Fatalf("Error: validTime24() method : expected t3 be valid. t3 was: %v", t3.String())
	}
	if t4.validTime24() {
		t.Fatalf("Error: validTime24() method : expected t4 be invalid. t4 was: %v", t4.String())
	}
}

func TestMinTime24(t *testing.T) {
	//test empty slice
	emptyList := []Time24{}
	minT1, err := minTime24(emptyList)
	tZero := Time24{0, 0, 0}
	if !equalsTime24(minT1, tZero) {
		t.Fatalf("Error: minTime24(): expected an zero Time24 return value. Returned value was: %v", minT1.String())
	}
	if err.Error() != "Error: the argument should not be empty slice" {
		t.Fatalf("Error: minTime24(): expected an error message as follow: \"Error: the argument should not be empty slice\"")
	}

	// testing non empty slice
	t1 := Time24{0, 1, 05}
	t2 := Time24{60, 01, 6}
	t3 := Time24{0000, 44, 5}
	t4 := Time24{0000, 200, 80}
	timeList := []Time24{t1, t2, t3, t4}

	minT2, err := minTime24(timeList)

	// expect t1 to be the minimum
	if !equalsTime24(minT2, t1) {
		t.Fatalf("Error: minTime24(): expected the minimum time to be 00:01:05. Returned value was: %v", minT2.String())
	}
	if err != nil {
		t.Fatalf("Error: minTime24(): expected a nil error message")
	}

	// testing non empty slice
	t5 := Time24{23, 1, 05}
	t6 := Time24{21, 41, 06}
	t7 := Time24{20, 44, 5}
	t8 := Time24{23, 59, 59}
	t9 := Time24{20, 44, 04}

	timeList2 := []Time24{t5, t6, t7, t8, t9}

	minT3, err := minTime24(timeList2)

	// expect t1 to be the minimum
	if !equalsTime24(minT3, t9) {
		t.Fatalf("Error: minTime24(): expected the minimum time to be 20:44:04. Returned value was: %v", minT3.String())
	}
	if err != nil {
		t.Fatalf("Error: minTime24(): expected a nil error message")
	}
}

//==========================================================================================
// Problem 4: use linear search to return the first index location of <x> in the slice <lst>
// return -1 if element not found
// works for int and string
// return -1 if type of <x> and <lst> is the same but not supported (i.e. not int or string)
// Panic if the type of <x> is not the same as the type of the elements in <slc>
//==========================================================================================

func TestLinearSearch(t *testing.T) {
	// sample data
	sampleIntArray := []int{4, 2, -1, 5, 0}
	sampleStringArray := []string{"cat", "nose", "egg"}
	sampleBoolArray := []bool{true, false, true}

	// test a valid int slice containing value
	returnedIdx := linearSearch(5, sampleIntArray)
	if returnedIdx != 3 {
		t.Fatalf("Error: linearSearch(): expected returned index to be 3 but it was: %v", returnedIdx)
	}

	// test a valid int slice containing value
	returnedIdx = linearSearch(0, sampleIntArray)
	if returnedIdx != 4 {
		t.Fatalf("Error: linearSearch(): expected returned index to be 4 but it was: %v", returnedIdx)
	}

	// test a valid int slice not containing value
	returnedIdx = linearSearch(3, sampleIntArray)
	if returnedIdx != -1 {
		t.Fatalf("Error: linearSearch(): expected returned index to be -1 but it was: %v", returnedIdx)
	}

	// test empty slice of int
	returnedIdx = linearSearch(3, []int{})
	if returnedIdx != -1 {
		t.Fatalf("Error: linearSearch(): expected returned index to be -1 but it was: %v", returnedIdx)
	}

	// test empty slice of string
	returnedIdx = linearSearch("egg", []string{})
	if returnedIdx != -1 {
		t.Fatalf("Error: linearSearch(): expected returned index to be -1 but it was: %v", returnedIdx)
	}

	// test valid slice of string containing element
	returnedIdx = linearSearch("egg", sampleStringArray)
	if returnedIdx != 2 {
		t.Fatalf("Error: linearSearch(): expected returned index to be 2 but it was: %v", returnedIdx)
	}

	// test valid slice of string not containing element
	returnedIdx = linearSearch("up", sampleStringArray)
	if returnedIdx != -1 {
		t.Fatalf("Error: linearSearch(): expected returned index to be -1 but it was: %v", returnedIdx)
	}

	// test valid slice of not supported type elements
	returnedIdx = linearSearch(true, sampleBoolArray)
	if returnedIdx != -1 {
		t.Fatalf("Error: linearSearch(): expected returned index to be -1 but it was: %v", returnedIdx)
	}

	linearSearch(3, sampleBoolArray)
	assert
}

//=====================================================
// Problem 5: return all the bit sequences of length n
//		      return empty slice if <n> 0 or negative
// Example:
// 		allBitSeqs(1) returns [[0] [1]].
// 		allBitSeqs(2) returns [[0 0] [0 1] [1 0] [1 1]]
//=====================================================
func TestAllBitSeqs(t *testing.T) {
	// test data
	emptySlice := make([][]int, 0)
	n1Slice := [][]int{{0}, {1}}
	n2Slice := [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	n3Slice := [][]int{{0, 0, 0}, {0, 0, 1}, {0, 1, 0}, {0, 1, 1}, {1, 0, 0}, {1, 0, 1}, {1, 1, 0}, {1, 1, 1}}

	// test 0 argument
	retSlice := allBitSeqs(0)
	if !reflect.DeepEqual(emptySlice, retSlice) {
		t.Fatalf("Error: allBitSeqs(): expected the returned slice to be [] but it was: %v", retSlice)
	}

	// test negative argument
	retSlice = allBitSeqs(-8)
	if !reflect.DeepEqual(emptySlice, retSlice) {
		t.Fatalf("Error: allBitSeqs(): expected the returned slice to be [] but it was: %v", retSlice)
	}

	// test argument being 1
	retSlice = allBitSeqs(1)
	if !reflect.DeepEqual(n1Slice, retSlice) {
		t.Fatalf("Error: allBitSeqs(): expected the returned slice to be [[0] [1]] but it was: %v", retSlice)
	}

	// test agument being 2
	retSlice = allBitSeqs(2)
	if !reflect.DeepEqual(n2Slice, retSlice) {
		t.Fatalf("Error: allBitSeqs(): expected the returned slice to be [0 0] [0 1] [1 0] [1 1]] but it was: %v", retSlice)
	}

	// test agument being 3
	retSlice = allBitSeqs(3)
	if !reflect.DeepEqual(n3Slice, retSlice) {
		t.Fatalf("Error: allBitSeqs(): expected the returned slice to be [[0 0 0] [0 0 1] [0 1 0] [0 1 1] [1 0 0] [1 0 1] [1 1 0] [1 1 1]] but it was: %v", retSlice)
	}
}
