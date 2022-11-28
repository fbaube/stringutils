package stringutils_test

import (
	"fmt"
	SU "github.com/fbaube/stringutils"
)

func ExampleYn() {
	fmt.Println(SU.Yn(true))
	fmt.Println(SU.Yn(false))
	// Output:
	// Y
	// n
}
