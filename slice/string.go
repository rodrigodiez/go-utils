// Provide utils to work with slices
package slice

// Finds out whether an slice contains a given string
func ContainsString(needle string, haystack *[]string) bool {
	return IndexOfString(haystack, needle) != -1
}

// Finds out the first index of a given string within an slice
func IndexOfString(needle string, haystack *[]string) int {
	for i, value := range *haystack {
		if value == needle {
			return i
		}
	}

	return -1
}
