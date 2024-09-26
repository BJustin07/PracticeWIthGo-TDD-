package iteration

import (
	"fmt"
	"testing"
)

func TestRepeatWord(t *testing.T) {
	got := RepeatWord("Behbeh", 3)
	want := "BehbehBehbehBehbeh"
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}

}

func ExampleRepeatWord() {
	repeat := RepeatWord("Behbeh", 2)
	fmt.Println(repeat)
	// Output: BehbehBehbeh
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatWord("Behbeh", 2)
	}
}
