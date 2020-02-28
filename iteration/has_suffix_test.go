package iteration

import "testing"

func TestHasSuffix(t *testing.T) {
	t.Run("has the suffix", func(t *testing.T) {
		hasSuffix := HasSuffix("Taiwan", "an")
		expected := true
		if hasSuffix != expected {
			t.Errorf("expected %t but got %t", expected, hasSuffix)
		}
	})

	t.Run("do not have the suffix", func(t *testing.T) {
		hasSuffix := HasSuffix("Taiwan", "na")
		expected := false
		if hasSuffix != expected {
			t.Errorf("expected %t but got %t", expected, hasSuffix)
		}
	})
}