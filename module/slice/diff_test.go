package slice

import (
	"fmt"
	"testing"
)

func TestDiffStrings(t *testing.T) {
	t.Run("diff slice", func(t *testing.T) {
		columns := []string{
			"id", "name", "middle_name", "gender",
		}
		primaryColumns := []string{
			"id",
		}

		notPrimaryColumns := []string{
			"name", "middle_name", "gender",
		}
		got := DiffStrings(columns, primaryColumns)

		if fmt.Sprintf("%v", notPrimaryColumns) != fmt.Sprintf("%v", got) {
			t.Errorf("expected  %v, got %v", notPrimaryColumns, got)
		}
	})
}
