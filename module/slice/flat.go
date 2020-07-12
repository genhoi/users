package slice

func Flat(slice [][]interface{}) []interface{} {
	var result []interface{}
	for _, item := range slice {
		result = append(result, item...)
	}

	return result
}
