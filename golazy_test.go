package golazy

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicateStrings(t *testing.T) {
	var tests = []struct {
		input    []string
		expected []string
	}{
		{[]string{}, []string{}},
		{[]string{"1", "2", "3"}, []string{"1", "2", "3"}},
		{[]string{"1", "2", "3", "1", "2", "3", "1", "2", "3"}, []string{"1", "2", "3"}},
		{[]string{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a"}, []string{"a"}},
		{[]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"},
			[]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}},
	}

	for _, test := range tests {
		if output := RemoveDuplicateStrings(test.input); !EqStringTest(test.expected, output) {
			errorString := fmt.Sprintf("Test Failed: %s inputted, %v expected, received: %v", test.input, test.expected, output)
			t.Error(errorString)
		}
	}
}

func TestRemoveDuplicateInts(t *testing.T) {
	var tests = []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 3, 1, 2, 3, 1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}},
	}

	for _, test := range tests {
		if output := RemoveDuplicateInts(test.input); !EqIntTest(test.expected, output) {
			errorString := fmt.Sprintf("Test Failed: %d inputted, %v expected, received: %v", test.input, test.expected, output)
			t.Error(errorString)
		}
	}
}

func TestRemoveDuplicateFloats(t *testing.T) {
	var tests = []struct {
		input    []float64
		expected []float64
	}{
		{[]float64{}, []float64{}},
		{[]float64{1.1, 2.2, 3.3}, []float64{1.1, 2.2, 3.3}},
		{[]float64{1, 2, 3, 1, 2, 3, 1, 2, 3}, []float64{1, 2, 3}},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}},
	}

	for _, test := range tests {
		if output := RemoveDuplicateFloats(test.input); !EqFloatTest(test.expected, output) {
			errorString := fmt.Sprintf("Test Failed: %f inputted, %v expected, received: %v", test.input, test.expected, output)
			t.Error(errorString)
		}
	}
}

func TestGetHost(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"http://ciao.com", "ciao.com"},
		{"https://google.com", "google.com"},
		{"http:google.com", ""},
		{"http//google.com", ""},
		{"//google.com", "google.com"},
	}

	for _, test := range tests {
		if output, _ := GetHost(test.input); test.expected != output {
			errorString := fmt.Sprintf("Test Failed: %s inputted, %v expected, received: %v", test.input, test.expected, output)
			t.Error(errorString)
		}
	}
}

func TestGetProtocol(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"http://ciao.com", "http"},
		{"https://google.com", "https"},
		{"http:google.com", "http"},
		{"http//google.com", ""},
		{"//google.com", ""},
	}

	for _, test := range tests {
		if output, _ := GetProtocol(test.input); test.expected != output {
			errorString := fmt.Sprintf("Test Failed: %s inputted, %v expected, received: %v", test.input, test.expected, output)
			t.Error(errorString)
		}
	}
}

//EqStringTest : Test if two slices of strings are equal
func EqStringTest(a, b []string) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

//EqIntTest : Test if two slices of ints are equal
func EqIntTest(a, b []int) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

//EqFloatTest : Test if two slices of floats are equal
func EqFloatTest(a, b []float64) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
