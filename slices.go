package slices

import "reflect"

func contains(s interface{}, h interface{}) bool {
	arr := reflect.ValueOf(s)

	if arr.Kind() == reflect.Slice {
		for i, n := 0, arr.Len(); i < n; i++ {
			if arr.Index(i).Interface() == h {
				return true
			}
		}
	}
	return false
}

// ContainsString return true if given slices contain given string, otherwise return false
func ContainsString(s []string, h string) bool {
	return contains(s, h)
}
