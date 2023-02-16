<p align="center">
  <img src="https://github.com/edoardottt/images/blob/main/golazy/golazy.png"><br>
  <b align="center">🌴golazy🍸</b><br>
  <p align="center">Golang module exporting general purpose functions I get tired of rewriting every time</p>
</p>
<p align="center">
<a href="https://edoardoottavianelli.it">
	<img src="https://github.com/edoardottt/golazy/actions/workflows/go.yml/badge.svg" alt="workflows" />
</a>
<a href="https://pkg.go.dev/github.com/edoardottt/golazy"><img src="https://pkg.go.dev/badge/github.com/edoardottt/golazy.svg" alt="Go Reference"></a>
</p>

----------

```
go get github.com/edoardottt/golazy
```
<br>

| Name | Description | 
| ----- | ------ |
| [ScanInputStdin() []string](https://github.com/edoardottt/golazy/blob/main/golazy.go#L39) | It returns the array of elements taken as input on stdin. |
| [RemoveDuplicateStrings(strSlice []string) []string](https://github.com/edoardottt/golazy/blob/main/golazy.go#L55) | It removes duplicates from a slice of strings taken as input and returns the result |
| [RemoveDuplicateInts(intSlice []int) []int](https://github.com/edoardottt/golazy/blob/main/golazy.go#L69) | It removes duplicates from a slice of integers taken as input and returns the result |
| [RemoveDuplicateFloats(floatSlice []float64) []float64](https://github.com/edoardottt/golazy/blob/main/golazy.go#L83) | It removes duplicates from a slice of floats taken as input and returns the result | 
| [AppendOutputToTxt(output string, filename string)](https://github.com/edoardottt/golazy/blob/main/golazy.go#L97) | It tries to append the output string in the file "filename" taken as input. |
| [AppendOutputToTxtAndExit(output string, filename string)](https://github.com/edoardottt/golazy/blob/main/golazy.go#L111) | It appends the output string in the file "filename" taken as input, but if it encounters an error it logs it and exits. |
| [GetHost(input string) (string, error)](https://github.com/edoardottt/golazy/blob/main/golazy.go#L129) | It takes as input a string and tries to parse it as url, if it's a well formatted url it returns the host (the domain if you prefer) |
| [GetProtocol(input string) (string, error)](https://github.com/edoardottt/golazy/blob/main/golazy.go#L141) | It takes as input a string and tries to parse it as url, if it's a well formatted url it returns the protocol |
| [HasProtocol(input string) bool](https://github.com/edoardottt/golazy/blob/main/golazy.go#L152) | It takes as input a string and checks if it has a protocol ( like in a URI/URL) |
| [RemoveProtocol(input string) string](https://github.com/edoardottt/golazy/blob/main/golazy.go#L160) | It removes the protocol from the input string (something://...). If it's not present it returns the input |
| [RemovePort(input string) string](https://github.com/edoardottt/golazy/blob/main/golazy.go#L170) | It removes port from the input string. If it's not present it returns the input |
| [SameDomain(url1 string, url2 string) bool](https://github.com/edoardottt/golazy/blob/main/golazy.go#L179) | It checks if two urls have the same domain |
| [GetPath(input string) (string, error)](https://github.com/edoardottt/golazy/blob/main/golazy.go#L196) | It returns the path of the input string (if correctly URL-formatted) |
| [ReadFileLineByLine(inputFile string) []string](https://github.com/edoardottt/golazy/blob/main/golazy.go#L206) | It reads from a file taken as input and returns a slice of strings (duplicates allowed). |
| [GenerateRandomUserAgent() string](https://github.com/edoardottt/golazy/blob/main/golazy.go#L286) | It generates a Random User Agent. |

<br>

License 📝
-------

This repository is under [GNU General Public License v3.0](https://github.com/edoardottt/golazy/blob/main/LICENSE).  
[edoardoottavianelli.it](https://www.edoardoottavianelli.it) to contact me.
