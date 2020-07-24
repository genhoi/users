package slice

func Flat(slice [][]interface{}) []interface{} {
	result := make([]interface{}, 0, len(slice))
	for _, item := range slice {
		result = append(result, item...)
	}

	return result
}
