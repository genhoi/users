package slice

import (
	"fmt"
	"testing"
)

func TestFlat(t *testing.T) {
	t.Run("flatten slice", func(t *testing.T) {
		v := [][]interface{}{
			{1, 2}, {3, 4}, {5, 6}, {7, 8},
		}

		expected := []interface{}{
			1, 2, 3, 4, 5, 6, 7, 8,
		}
		got := Flat(v)

		if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", got) {
			t.Errorf("expected  %v, got %v", expected, got)
		}
	})
	t.Run("flatten of 2 slices", func(t *testing.T) {
		v := [][]interface{}{
			{
				[]interface{}{1, 2},
				[]interface{}{3, 4},
			},
			{
				[]interface{}{5, 6},
				[]interface{}{7, 8},
			},
		}

		expected := []interface{}{
			[]interface{}{1, 2},
			[]interface{}{3, 4},
			[]interface{}{5, 6},
			[]interface{}{7, 8},
		}
		got := Flat(v)

		if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", got) {
			t.Errorf("expected  %v, got %v", expected, got)
		}
	})
}

func BenchmarkFlat(b *testing.B) {
	v := [][]interface{}{
		{
			[]interface{}{1, 2},
			[]interface{}{3, 4},
		},
		{
			[]interface{}{5, 6},
			[]interface{}{7, 8},
		},
	}

	for i := 0; i < b.N; i++ {
		_ = Flat(v)
	}
}
