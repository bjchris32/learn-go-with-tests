package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertMessage := func(t *testing.T, got string, expected string) {
		t.Helper() // for tracing the failure line number in the caller of this assertion
		if got != expected {
			t.Errorf("got %q expected %q", got, expected)
		}
	}

	t.Run("repeat the string for 8 times when number of iterations is given", func(t *testing.T) {
		got := Repeat("a", 8)
		expected := "aaaaaaaa"
		assertMessage(t, got, expected)
	})

	t.Run("repeat the string for 1 times when there is no iteration parameter given", func(t *testing.T) {
		got := Repeat("a")
		expected := "a"
		assertMessage(t, got, expected)
	})
}

func ExampleRepeat_1() {
	repeated := Repeat("ab", 2)
	fmt.Println(repeated)
	// Output: abab
}

func ExampleRepeat_2() {
	repeated := Repeat("ab", -1)
	fmt.Println(repeated)
	// Output: ab
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 2)
	}
}
