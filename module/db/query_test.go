package db

import (
	"github.com/genhoi/users/module/slice"
	"testing"
)

func TestDeduplicateItemsByPrimaryKeys(t *testing.T) {
	t.Run("deduplicate", func(t *testing.T) {
		type TestEntity struct {
			Name string `db:"name, primarykey"`
			Id   int    `db:"id, primarykey"`
			Text string `db:"text"`
		}

		entities := []interface{}{
			TestEntity{"one", 1, "Test one 1"},
			TestEntity{"two", 2, "Test two 1"},
			TestEntity{"one", 1, "Test one 2"},
		}

		got, _ := deduplicateItemsByPrimaryKeys(entities)
		want := []interface{}{
			TestEntity{"one", 1, "Test one 2"},
			TestEntity{"two", 2, "Test two 1"},
		}

		if !slice.DeepEqual(want, got) {
			t.Errorf("got %v, want  %v", got, want)
		}
	})
	t.Run("deduplicate with only primary keys", func(t *testing.T) {
		type TestEntity struct {
			Name string `db:"name, primarykey"`
			Id   int    `db:"id, primarykey"`
		}

		entities := []interface{}{
			TestEntity{"one", 1},
			TestEntity{"two", 2},
			TestEntity{"one", 1},
		}

		got, _ := deduplicateItemsByPrimaryKeys(entities)
		want := []interface{}{
			TestEntity{"one", 1},
			TestEntity{"two", 2},
		}

		if !slice.DeepEqual(want, got) {
			t.Errorf("got %v, want  %v", got, want)
		}
	})
	t.Run("deduplicate by one primary key", func(t *testing.T) {
		type TestEntity struct {
			Name string `db:"name"`
			Id   int    `db:"id, primarykey"`
		}

		entities := []interface{}{
			TestEntity{"one", 1},
			TestEntity{"two", 2},
			TestEntity{"three", 1},
		}

		got, _ := deduplicateItemsByPrimaryKeys(entities)
		want := []interface{}{
			TestEntity{"three", 1},
			TestEntity{"two", 2},
		}

		if !slice.DeepEqual(want, got) {
			t.Errorf("got %v, want  %v", got, want)
		}
	})
	t.Run("deduplicate ptrs", func(t *testing.T) {
		type TestEntity struct {
			Name string `db:"name, primarykey"`
			Id   int    `db:"id, primarykey"`
			Text string `db:"text"`
		}

		entities := []interface{}{
			&TestEntity{"one", 1, "Test one 1"},
			&TestEntity{"two", 2, "Test two 1"},
			&TestEntity{"one", 1, "Test one 2"},
		}

		got, _ := deduplicateItemsByPrimaryKeys(entities)
		want := []interface{}{
			&TestEntity{"one", 1, "Test one 2"},
			&TestEntity{"two", 2, "Test two 1"},
		}

		if !slice.DeepEqual(want, got) {
			t.Errorf("got %v, want  %v", got, want)
		}
	})
}
