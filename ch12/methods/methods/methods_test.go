package methods_test

import (
	"learning/methods"
	"strings"
)

//func ExamplePrintDuration() {
//	methods.Print(time.Hour)
//	// Output:
//}

func ExamplePrintReplacer() {
	methods.Print(new (strings.Replacer))
	// Output:
}

