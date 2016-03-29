// Provide utils to work with slices
package slice

// Finds out whether an slice contains a given string
func containsString(haystack *[]string, needle string) bool {
	return indexOfString(haystack, needle) != -1
}

// Finds out whether an slice contains a given string
func indexOfString(haystack *[]string, needle string) int {
	for i, value := range *haystack {
		if value == needle {
			return i
		}
	}

	return -1
}
