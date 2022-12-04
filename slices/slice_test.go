package slices

import "testing"

func TestUnique(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		t.Parallel()

		got := Unique([]int{1, 2, 3})
		if got == false {
			t.Fail()
		}
	})

	t.Run("strings", func(t *testing.T) {
		t.Parallel()

		got := Unique([]string{"foo", "bar"})
		if got == false {
			t.Fail()
		}
	})
}

func TestExists(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		got := Exists([]int{1, 2, 3}, 2)
		if got == false {
			t.Fail()
		}
	})

	t.Run("strings", func(t *testing.T) {
		got := Exists([]string{"foo", "bar"}, "bar")
		if got == false {
			t.Fail()
		}
	})
}
