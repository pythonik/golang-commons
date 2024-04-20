package arrayutils

import "testing"

func TestIsEmpty(t *testing.T) {
	t.Run("Empty slice", func(t *testing.T) {
		s := []int{}
		if !IsEmpty(s) {
			t.Errorf("IsEmpty(%v) = false; want true", s)
		}
	})
	t.Run("Non-empty slice", func(t *testing.T) {
		s := []int{1, 2, 3}
		if IsEmpty(s) {
			t.Errorf("IsEmpty(%v) = true; want false", s)
		}
	})
	t.Run("nil slice", func(t *testing.T) {
		var s []int
		if IsEmpty(s) {
			t.Errorf("IsEmpty(%v) = false; want true", s)
		}
	})
}
