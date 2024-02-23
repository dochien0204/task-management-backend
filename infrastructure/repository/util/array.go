package util

func InArray[T comparable](element T, array []T) bool {
	for _, value := range array {
		if element == value {
			return true
		}
	}

	return false
}

func RemoveItem(str string, array []string) []string {
	result := []string{}
	for _, element := range array {
		if element != str {
			result = append(result, element)
		}
	}

	return result
}
