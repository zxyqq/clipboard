package clipboard_test

import (
	"fmt"

	"github.com/zxyqq/clipboard"
)

func Example() {
	clipboard.WriteAll("Привет, мир!")
	text, _ := clipboard.ReadAll()
	fmt.Println(text)

	// Output:
	// Привет, мир!
}
