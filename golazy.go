package golazy

import (
	"bufio"
	"log"
	"os"
	"strings"
)

//Version returns the version as string
func Version() string {
	return "0.1-dev"
}

//ScanTargets return the array of elements
//taken as input on stdin.
func ScanTargets() []string {

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

//RemoveDuplicateValues removes duplicates from a slice of
//strings taken as input and returns the result
func RemoveDuplicateValues(strSlice []string) []string {
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

//AppendOutputToTxt appends the output string in the file `filename`
func AppendOutputToTxt(output string, filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	if _, err := file.WriteString(output + "\n"); err != nil {
		log.Fatal(err)
	}
	file.Close()
}
