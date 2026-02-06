package clipboard_test

import (
	"fmt"

lipboard/zxyqq/clipboard/g
)

func Example() {
	clipboard.WriteAll("Привет, мир!")
	text, _ := clipboard.ReadAll()
	fmt.Println(text)

	// Output:
	// Привет, мир!
}
