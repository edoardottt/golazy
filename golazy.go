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
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	Permission0644 = 0644
)

// Version returns the version as string.
func Version() string {
	return "0.1.4"
}

// ScanInputStdin return the array of elements
// taken as input on stdin.
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

// RemoveDuplicateValues removes duplicates from a slice of
// T taken as input and returns the result.
func RemoveDuplicateValues[T comparable](slice []T) []T {
	keys := make(map[T]bool)
	list := []T{}

	for _, entry := range slice {
		if ok := keys[entry]; !ok {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

// AppendOutputToTxt tries to append the output string in the file `filename`
// taken as input.
func AppendOutputToTxt(output string, filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, Permission0644)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := file.WriteString(output + "\n"); err != nil {
		log.Fatal(err)
	}

	file.Close()
}

// AppendOutputToTxtAndExit appends the output string in the file `filename`
// taken as input.
// ---> If it encounters an error, it logs the error and exits !!!!
func AppendOutputToTxtAndExit(output string, filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, Permission0644)
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

// GetHost takes as input a string and
// tries to parse it as url, if it's a
// well formatted url this function returns
// the host (the domain if you prefer).
func GetHost(input string) (string, error) {
	u, err := url.Parse(input)
	if err != nil {
		return "", err
	}

	return u.Host, nil
}

// GetProtocol takes as input a string and
// tries to parse it as url, if it's a
// well formatted url this function returns
// the protocol (the scheme if you prefer).
func GetProtocol(input string) (string, error) {
	u, err := url.Parse(input)
	if err != nil {
		return "", err
	}

	return u.Scheme, nil
}

// HasProtocol takes as input a string and
// checks if it has a protocol ( like in a
// URI/URL).
func HasProtocol(input string) bool {
	u, err := url.Parse(input)
	if err != nil {
		return false
	}

	return u.Scheme != ""
}

// RemoveProtocol removes the protocol from
// the input string (something://...)
// If it's not present it returns the input.
func RemoveProtocol(input string) string {
	if HasProtocol(input) {
		res := strings.Index(input, "://")
		if res >= 0 {
			return input[res+3:]
		}
	}

	return input
}

// RemovePort removes port from the input string
// If it's not present it returns the input.
func RemovePort(input string) string {
	res := strings.Index(input, ":")
	if res >= 0 {
		return input[:res]
	}

	return input
}

// SameDomain checks if two urls have the same domain.
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

// GetPath returns the path of the input string
// (if correctly URL-formatted).
func GetPath(input string) (string, error) {
	u, err := url.Parse(input)
	if err != nil {
		return "", err
	}

	return u.Path, nil
}

// ReadFileLineByLine reads from a file taken as input
// and returns a slice of strings (duplicates allowed).
func ReadFileLineByLine(inputFile string) []string {
	var text []string

	file, err := os.Open(inputFile)

	if err != nil {
		log.Fatalf("failed to open %s ", inputFile)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		dir := scanner.Text()
		if len(dir) > 0 {
			text = append(text, dir)
		}
	}

	file.Close()

	return text
}

// genOsString generates a random OS string for a User Agent.
func genOsString() string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	// Operating system.
	var OsStrings = []string{
		"Macintosh; Intel Mac OS X 10_10",
		"Windows NT 10.0",
		"Windows NT 5.1",
		"Windows NT 6.1; WOW64",
		"Windows NT 6.1; Win64; x64",
		"X11; Linux x86_64",
	}

	return OsStrings[rng.Intn(len(OsStrings))]
}

// genFirefoxUA generates a random Firefox User Agent.
func genFirefoxUA() string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	// Firefox versions.
	var FirefoxVersions = []float32{
		122.0,
		121.0,
		120.0,
		119.0,
		118.0,
		117.0,
		116.0,
		115.0,
		114.0,
		113.0,
		112.0,
		111.0,
		110.0,
		109.0,
		108.0,
		107.0,
		106.0,
	}

	version := FirefoxVersions[rng.Intn(len(FirefoxVersions))]

	return fmt.Sprintf("Mozilla/5.0 (%s; rv:%.1f) Gecko/20100101 Firefox/%.1f", genOsString(), version, version)
}

// genChromeUA generates a random Chrome User Agent.
func genChromeUA() string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	// Chrome versions.
	var ChromeVersions = []string{
		"107.0.5304",
		"106.0.5249",
		"105.0.5195",
		"104.0.5112",
		"103.0.5060",
		"102.0.5005",
		"101.0.4951",
		"100.0.4896",
		"99.0.4844",
		"98.0.4758",
		"97.0.4692",
		"96.0.4664",
		"95.0.4638",
		"94.0.4606",
		"93.0.4577",
		"92.0.4515",
		"91.0.4472",
		"90.0.4430",
	}

	version := ChromeVersions[rng.Intn(len(ChromeVersions))]

	return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36",
		genOsString(), version)
}

// GenerateRandomUserAgent generates a random user agent
// (can be Chrome or Firefox).
func GenerateRandomUserAgent() string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	decision := rng.Intn(100)

	var ua string
	if decision%2 == 0 {
		ua = genChromeUA()
	} else {
		ua = genFirefoxUA()
	}

	return ua
}
