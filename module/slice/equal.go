package slice

import "reflect"

func DeepEqual(x []interface{}, y []interface{}) bool {
	if len(x) != len(y) {
		return false
	}
	i := 0
	for _, yItem := range y {
		for _, xItem := range x {
			if reflect.DeepEqual(yItem, xItem) {
				i++
			}
		}
	}

	if i != len(x) {
		return false
	}

	return true
}
