package arrays_and_string_manipulation

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	secret := []byte(`cake pound steal`)

	reverseWordsInPlace(secret)
	fmt.Printf("%s", secret)
}
