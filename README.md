<p align="center">
  <img src="https://github.com/edoardottt/images/blob/main/golazy/golazy.png"><br>
  <b align="center">🌴golazy🍸</b><br>
  <p align="center">A Go module exporting general purpose functions I get tired of rewriting every time</p>
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

<br>

License 📝
-------

This repository is under [GNU General Public License v3.0](https://github.com/edoardottt/golazy/blob/main/LICENSE).  
[edoardoottavianelli.it](https://www.edoardoottavianelli.it) to contact me.
