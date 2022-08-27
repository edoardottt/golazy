/*
=======================
golazy
=======================
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.
You should have received a copy of the GNU General Public License
along with this program.  If not, see http://www.gnu.org/licenses/.
	@Repository:  https://github.com/edoardottt/golazy
	@Author:      edoardottt, https://www.edoardoottavianelli.it
*/

package golazy_test

import (
	"fmt"
	"testing"

	"github.com/edoardottt/golazy"
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
		{[]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
			"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
			"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
			"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"},
			[]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
				"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}},
	}

	for _, test := range tests {
		if output := golazy.RemoveDuplicateStrings(test.input); !EqStringTest(test.expected, output) {
			errorString := fmt.Sprintf("Test Failed: %s inputted, %v expected, received: %v",
				test.input, test.expected, output)
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
		if output := golazy.RemoveDuplicateInts(test.input); !EqIntTest(test.expected, output) {
			errorString := fmt.Sprintf("Test Failed: %d inputted, %v expected, received: %v",
				test.input, test.expected, output)
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
		if output := golazy.RemoveDuplicateFloats(test.input); !EqFloatTest(test.expected, output) {
			errorString := fmt.Sprintf("Test Failed: %f inputted, %v expected, received: %v",
				test.input, test.expected, output)
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
		if output, _ := golazy.GetHost(test.input); test.expected != output {
			errorString := fmt.Sprintf("Test Failed: %s inputted, %v expected, received: %v",
				test.input, test.expected, output)
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
		if output, _ := golazy.GetProtocol(test.input); test.expected != output {
			errorString := fmt.Sprintf("Test Failed: %s inputted, %v expected, received: %v",
				test.input, test.expected, output)
			t.Error(errorString)
		}
	}
}

func TestHasProtocol(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"", false},
		{"http://ciao.com", true},
		{"https://google.com", true},
		{"http:google.com", false},
		{"http//google.com", false},
		{"//google.com", false},
	}

	for _, test := range tests {
		if output := golazy.HasProtocol(test.input); test.expected != output {
			errorString := fmt.Sprintf("Test Failed: %s inputted, %v expected, received: %v",
				test.input, test.expected, output)
			t.Error(errorString)
		}
	}
}

func TestRemoveProtocol(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"http://ciao.com/ciao/?ciao=1", "ciao.com/ciao/?ciao=1"},
		{"https://google.com/#heyyyyyy", "google.com/#heyyyyyy"},
		{"http:google.com", "http:google.com"},
		{"http//google.com", "http//google.com"},
		{"//google.com", "//google.com"},
		{"https://subdomain.domain.tld/path1/path2?ciao=paramvalue4", "subdomain.domain.tld/path1/path2?ciao=paramvalue4"},
	}

	for _, test := range tests {
		if output := golazy.RemoveProtocol(test.input); test.expected != output {
			errorString := fmt.Sprintf("Test Failed: %s inputted, %v expected, received: %v",
				test.input, test.expected, output)
			t.Error(errorString)
		}
	}
}

func TestRemovePort(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"subdomain.google.com", "subdomain.google.com"},
		{"subdomain.google.com:443", "subdomain.google.com"},
	}

	for _, test := range tests {
		if output := golazy.RemovePort(test.input); test.expected != output {
			errorString := fmt.Sprintf("Test Failed: %s inputted, %v expected, received: %v",
				test.input, test.expected, output)
			t.Error(errorString)
		}
	}
}

func TestSameDomain(t *testing.T) {
	var tests = []struct {
		input1   string
		input2   string
		expected bool
	}{
		{"", "", false},
		{"http://subdomain.google.com", "http://subdomain.google.com", true},
		{"http://subdomain.google.com", "https://subdomain.google.com", true},
		{"http://subdomain.googla.com", "https://subdomain.google.com", false},
		{"http://subdomain.google.com:80", "https://subdomain.google.com", false},
		{"http://subdomain.google.com.com", "https://subdomain.google.com", false},
	}

	for _, test := range tests {
		if output := golazy.SameDomain(test.input1, test.input2); test.expected != output {
			errorString := fmt.Sprintf("Test Failed: %s and %s inputted, %v expected, received: %v",
				test.input1, test.input2, test.expected, output)
			t.Error(errorString)
		}
	}
}

func TestGetPath(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"https://subdomain.google.com", ""},
		{"https://subdomain.google.com:443/ciao", "/ciao"},
		{"https://subdomain.google.com:443/ciao/", "/ciao/"},
		{"https://subdomain.google.com:443/ciao/#", "/ciao/"},
		{"https://subdomain.google.com:443/ciao/?q=1", "/ciao/"},
	}

	for _, test := range tests {
		if output, _ := golazy.GetPath(test.input); test.expected != output {
			errorString := fmt.Sprintf("Test Failed: %s inputted, %v expected, received: %v",
				test.input, test.expected, output)
			t.Error(errorString)
		}
	}
}

// EqStringTest : Test if two slices of strings are equal.
func EqStringTest(input1, input2 []string) bool {
	// If one is nil, the other must also be nil.
	if (input1 == nil) != (input2 == nil) {
		return false
	}

	if len(input1) != len(input2) {
		return false
	}

	for i := range input1 {
		if input1[i] != input2[i] {
			return false
		}
	}

	return true
}

// EqIntTest : Test if two slices of ints are equal.
func EqIntTest(input1, input2 []int) bool {
	// If one is nil, the other must also be nil.
	if (input1 == nil) != (input2 == nil) {
		return false
	}

	if len(input1) != len(input2) {
		return false
	}

	for i := range input1 {
		if input1[i] != input2[i] {
			return false
		}
	}

	return true
}

// EqFloatTest : Test if two slices of floats are equal.
func EqFloatTest(input1, input2 []float64) bool {
	// If one is nil, the other must also be nil.
	if (input1 == nil) != (input2 == nil) {
		return false
	}

	if len(input1) != len(input2) {
		return false
	}

	for i := range input1 {
		if input1[i] != input2[i] {
			return false
		}
	}

	return true
}
