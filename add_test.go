package add

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Fatalf("The result is not 5.")
		return
	}
}