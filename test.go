package test

import (
	"fmt"

	golazy "github.com/edoardottt/golazy/pkg"
)

func test1() {
	fmt.Println(golazy.RemoveDuplicateValues([]string{"golazy"}))
}
