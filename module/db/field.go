package db

import (
	"reflect"
	"strconv"
	"strings"
	"sync"
)

var (
	tagCache   = map[reflect.Type][]fieldTag{}
	tagCacheMx = sync.RWMutex{}
)

type fieldTag struct {
	number     int
	name       string
	primaryKey bool
	nullable   bool
}

type field struct {
	tag   fieldTag
	value interface{}
}

func fields(entity interface{}) []field {
	v := reflect.ValueOf(entity)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	tags := getTags(v)
	f := make([]field, len(tags))
	for i, tag := range tags {
		value := v.Field(tag.number).Interface()
		f[i] = field{tag, value}
	}

	return f
}

func getTags(v reflect.Value) []fieldTag {
	reflectType := v.Type()

	tagCacheMx.RLock()
	cache, ok := tagCache[reflectType]
	tagCacheMx.RUnlock()
	if ok {
		return cache
	}

	var (
		tags []fieldTag
		tag  string

		columnName string
		primaryKey bool
		nullable   bool
	)

	numField := v.NumField()
	for i := 0; i < numField; i++ {
		structField := reflectType.Field(i)
		tag = structField.Tag.Get("db")
		if tag == "" {
			continue
		}
		columnName, primaryKey, nullable = parseTag(tag)
		tags = append(tags, fieldTag{
			number:     i,
			name:       columnName,
			primaryKey: primaryKey,
			nullable:   nullable,
		})
	}

	tagCacheMx.Lock()
	tagCache[reflectType] = tags
	tagCacheMx.Unlock()

	return tags
}

func parseTag(tag string) (columnName string, primaryKey bool, nullable bool) {
	arguments := strings.Split(tag, ",")

	primaryKey, nullable = false, false
	columnName = arguments[0]

	for _, argString := range arguments[1:] {
		argString = strings.TrimSpace(argString)
		arg := strings.SplitN(argString, ":", 2)

		switch arg[0] {
		case "primarykey":
			primaryKey = true
		case "nullable":
			nullable = true
		}
	}

	return columnName, primaryKey, nullable
}

func Binding(entity interface{}, start int) string {
	names := ExtractNames(entity)

	binding := make([]string, len(names))
	i := start
	for k := range names {
		binding[k] = "$" + strconv.Itoa(i)
		i++
	}

	return "(" + strings.Join(binding, ", ") + ")"
}

func ExtractNames(entity interface{}) []string {
	fields := fields(entity)
	names := make([]string, 0, len(fields))

	for _, item := range fields {
		names = append(names, item.tag.name)
	}

	return names
}

func ExtractValues(entity interface{}) []interface{} {
	fields := fields(entity)
	values := make([]interface{}, len(fields))
	for i, item := range fields {
		values[i] = item.value
	}

	return values
}

func ExtractPrimaryKeys(entity interface{}) []string {
	fields := fields(entity)
	names := make([]string, 0, len(fields))

	for _, item := range fields {
		if item.tag.primaryKey {
			names = append(names, item.tag.name)
		}
	}

	return names
}

func ExtractPrimaryKeysValues(entity interface{}) []interface{} {
	fields := fields(entity)
	values := make([]interface{}, 0, len(fields))

	for _, item := range fields {
		if item.tag.primaryKey {
			values = append(values, item.value)
		}
	}

	return values
}
