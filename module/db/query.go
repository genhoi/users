package db

import (
	"errors"
	"fmt"
	"github.com/genhoi/users/module/slice"
	"strings"
)

type Query struct {
	tableName string
}

func NewQuery(tableName string) *Query {
	return &Query{tableName}
}

// ViaUpsert: inserts data if conflicting row hasn't been found, else it will update an existing one
func (q *Query) Upsert(items []interface{}) (sql string, values []interface{}, err error) {
	count := len(items)
	if count == 0 {
		err := errors.New("slice items should contain 1 elements or more")
		return "", nil, err
	}

	columns, primaryColumns, values, bindings, err := prepareBatchInsert(items)
	if err != nil {
		return "", nil, err
	}

	notPrimaryColumns := slice.DiffStrings(columns, primaryColumns)
	excludedColumns := make([]string, len(notPrimaryColumns))
	for i, v := range notPrimaryColumns {
		excludedColumns[i] = v + " = EXCLUDED." + v
	}

	data := struct {
		TableName string
		Columns   string
		Bindings  string
		Conflict  string
		UpdateSet string
	}{
		q.tableName,
		"(" + strings.Join(columns, ", ") + ")",
		strings.Join(bindings, ", "),
		"(" + strings.Join(primaryColumns, ", ") + ")",
		strings.Join(excludedColumns, ", "),
	}

	sql = fmt.Sprintf(
		`INSERT INTO %s %s VALUES %s ON CONFLICT %s DO UPDATE SET %s;`,
		data.TableName,
		data.Columns,
		data.Bindings,
		data.Conflict,
		data.UpdateSet,
	)

	return sql, values, nil
}

func prepareBatchInsert(items []interface{}) (
	columns []string,
	primaryColumns []string,
	values []interface{},
	bindings []string,
	err error,
) {
	if len(items) == 0 {
		err = errors.New("slice items should contain 1 elements or more")
		return
	}

	columns = ExtractNames(items[0])
	primaryColumns = ExtractPrimaryKeys(items[0])

	// Ensure that no rows proposed for insertion within the same command have duplicate constrained values
	// ----------------------------------------------------------------------------------------------------
	uniqueItems, err := deduplicateItemsByPrimaryKeys(items)
	if err != nil {
		return
	}
	// ----------------------------------------------------------------------------------------------------
	count := len(uniqueItems)
	itemsValues := make([][]interface{}, count)
	bindings = make([]string, count)

	i := 1
	for k, v := range uniqueItems {
		itemsValues[k] = ExtractValues(v)
		bindings[k] = Binding(v, i)
		i += len(columns)
	}

	values = slice.Flat(itemsValues)
	return
}

func deduplicateItemsByPrimaryKeys(items []interface{}) ([]interface{}, error) {
	count := len(items)
	if count == 0 {
		err := errors.New("slice items should contain 1 elements or more")
		return nil, err
	}

	mapItemsByPrimaryKeys := make(map[string]interface{})

	for _, v := range items {
		keyValues := ExtractPrimaryKeysValues(v)
		keyValuesAsString := make([]string, 0, len(keyValues))

		for _, keyValue := range keyValues {
			keyValuesAsString = append(keyValuesAsString, fmt.Sprintf("%v", keyValue))
		}
		key := strings.Join(keyValuesAsString, "_")
		mapItemsByPrimaryKeys[key] = v
	}

	uniqueItems := make([]interface{}, 0, len(mapItemsByPrimaryKeys))
	for _, item := range mapItemsByPrimaryKeys {
		uniqueItems = append(uniqueItems, item)
	}

	return uniqueItems, nil
}
