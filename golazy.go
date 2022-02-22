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

package golazy

import (
	"bufio"
	"log"
	"net/url"
	"os"
	"strings"
)

//Version returns the version as string
func Version() string {
	return "0.1-dev"
}

//ScanInputStdin return the array of elements
//taken as input on stdin.
func ScanInputStdin() []string {

	var result []string
	// accept domains on stdin
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		if strings.TrimSpace(sc.Text()) != "" {
			domain := strings.ToLower(sc.Text())
			result = append(result, domain)
		}
	}
	return result
}

//RemoveDuplicateStrings removes duplicates from a slice of
//strings taken as input and returns the result
func RemoveDuplicateStrings(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

//RemoveDuplicateInts removes duplicates from a slice of
//integers taken as input and returns the result
func RemoveDuplicateInts(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

//RemoveDuplicateFloats removes duplicates from a slice of
//floats taken as input and returns the result
func RemoveDuplicateFloats(floatSlice []float64) []float64 {
	keys := make(map[float64]bool)
	list := []float64{}
	for _, entry := range floatSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

//AppendOutputToTxt tries to append the output string in the file `filename`
//taken as input.
func AppendOutputToTxt(output string, filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := file.WriteString(output + "\n"); err != nil {
		log.Fatal(err)
	}
	file.Close()
}

//AppendOutputToTxtAndExit appends the output string in the file `filename`
//taken as input.
//---> If it encounters an error, it logs the error and exits !!!!
func AppendOutputToTxtAndExit(output string, filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if _, err := file.WriteString(output + "\n"); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	file.Close()
}

//GetHost takes as input a string and
//tries to parse it as url, if it's a
//well formatted url this function returns
//the host (the domain if you prefer)
func GetHost(input string) (string, error) {
	u, err := url.Parse(input)
	if err != nil {
		return "", err
	}
	return u.Host, nil
}

//GetProtocol takes as input a string and
//tries to parse it as url, if it's a
//well formatted url this function returns
//the protocol (the scheme if you prefer)
func GetProtocol(input string) (string, error) {
	u, err := url.Parse(input)
	if err != nil {
		return "", err
	}
	return u.Scheme, nil
}

//HasProtocol takes as input a string and
//checks if it has a protocol ( like in a
//URI/URL)
func HasProtocol(input string) bool {
	res := strings.Index(input, "://")
	return res >= 0
}

//RemoveProtocol removes the protocol from
//the input string (something://...)
//If it's not present it returns the input
func RemoveProtocol(input string) string {
	res := strings.Index(input, "://")
	if res >= 0 {
		return input[res+3:]
	}
	return input
}

//RemovePort removes port from the input string
//If it's not present it returns the input
func RemovePort(input string) string {
	res := strings.Index(input, ":")
	if res >= 0 {
		return input[:res]
	}
	return input
}

//SameDomain checks if two urls have the same domain
func SameDomain(url1 string, url2 string) bool {
	u1, err := url.Parse(url1)
	if err != nil {
		return false
	}
	u2, err := url.Parse(url2)
	if err != nil {
		return false
	}
	if u1.Host == "" || u2.Host == "" {
		return false
	}
	return u1.Host == u2.Host
}

//GetPath returns the path of the input string
//(if correctly URL-formatted)
func GetPath(input string) (string, error) {
	u, err := url.Parse(input)
	if err != nil {
		return "", err
	}
	return u.Path, nil
}

//ReadFileLineByLine read from a file taken as input
//and returns a slice of strings (duplicates allowed).
func ReadFileLineByLine(inputFile string) []string {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("failed to open %s ", inputFile)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	var dir = ""
	for scanner.Scan() {
		dir = scanner.Text()
		if len(dir) > 0 {
			text = append(text, dir)
		}
	}
	file.Close()
	return text
}
