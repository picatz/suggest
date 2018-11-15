package suggest

import (
	"fmt"
	"testing"
)

func TestSuggest(t *testing.T) {
	query := "apples and oranges"

	suggestions, err := Suggest(query)

	if err != nil {
		t.Error(err)
	}

	if len(suggestions) <= 0 {
		t.Error("expected suggestions, got none")
	}
}

func ExampleSuggest() {
	query := "apples and oranges"

	suggestions, err := Suggest(query)

	if err != nil {
		panic(err)
	}

	count := len(suggestions)

	fmt.Println(count > 0)

	// Output: true
}
