package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sumAdd := Add(2, 2)
	expected := 4
	if sumAdd != expected {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sumAdd, expected)
	}
}

func ExampleAdd() {
	sum := Add(34, 35)
	fmt.Println(sum)
	// Output: 69
}
