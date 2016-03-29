// Provide utils to work with slices
package slice

// Finds out whether an slice contains a given string
func ContainsString(haystack *[]string, needle string) bool {
	return IndexOfString(haystack, needle) != -1
}

// Finds out whether an slice contains a given string
func IndexOfString(haystack *[]string, needle string) int {
	for i, value := range *haystack {
		if value == needle {
			return i
		}
	}

	return -1
}
