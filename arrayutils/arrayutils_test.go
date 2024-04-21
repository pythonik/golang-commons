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
		if !IsEmpty(s) {
			t.Errorf("IsEmpty(%v) = false; want true", s)
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		var s []int
		if IsNotEmpty(s) {
			t.Errorf("IsEmpty(%v) = true; want false", s)
		}
	})
}

func TestContains(t *testing.T) {
	s := []int{1, 2, 3}
	t.Run("Element is in slice", func(t *testing.T) {
		if !Contains(s, 2) {
			t.Errorf("Contains(%v, 2) = false; want true", s)
		}
	})
	t.Run("Element is not in slice", func(t *testing.T) {
		if Contains(s, 4) {
			t.Errorf("Contains(%v, 4) = true; want false", s)
		}
	})
	t.Run("Element is not in slice when slice is nil", func(t *testing.T) {
		if Contains(s, 4) {
			t.Errorf("Contains(%v, 4) = true; want false", s)
		}
	})
}

func TestReverse(t *testing.T) {
	t.Run("Reverse int slice", func(t *testing.T) {
		s := []int{1, 2, 3}
		Reverse(s)
		want := []int{3, 2, 1}
		for i := range s {
			if s[i] != want[i] {
				t.Errorf("Reverse(%v) = %v; want %v", s, s, want)
				break
			}
		}
	})
	t.Run("Reverse string slice", func(t *testing.T) {
		s := []string{"1", "2", "3"}
		Reverse(s)
		want := []string{"3", "2", "1"}
		for i := range s {
			if s[i] != want[i] {
				t.Errorf("Reverse(%v) = %v; want %v", s, s, want)
				break
			}
		}
	})
}

func TestHead(t *testing.T) {
	t.Run("Head of int slice", func(t *testing.T) {
		s := []int{1, 2, 3}
		got := Head(s)
		want := 1
		if got.IsPresent() && got.Value() != want {
			t.Errorf("Head(%v) = %v; want %v", s, got, want)
		}
	})
	t.Run("Head of string slice", func(t *testing.T) {
		s := []string{"1", "2", "3"}
		got := Head(s)
		want := "1"
		if got.IsPresent() && got.Value() != want {
			t.Errorf("Head(%v) = %v; want %v", s, got, want)
		}
	})
}
