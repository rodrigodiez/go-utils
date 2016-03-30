// Provide utils to work with slices
package slice

// Finds out whether an slice contains a given string
//
//	# true
//	slice.ContainsString("Foo", &string[]{"Foo", "Bar"})
//	# false
//	slice.ContainsString("Baz", &string[]{"Foo", "Bar"})
func ContainsString(needle string, haystack *[]string) bool {
	return IndexOfString(needle, haystack) != -1
}

// Finds out the first index of a given string within an slice
//
//	# 0
//	slice.IndexOfString("Foo", &string[]{"Foo", "Bar"})
//	# 1
//	slice.IndexOfString("Foo", &string[]{"Foo", "Foo"})
//	# -1
//	slice.IndexOfString("Foo", &string[]{"Bar"})
func IndexOfString(needle string, haystack *[]string) int {
	for i, value := range *haystack {
		if value == needle {
			return i
		}
	}

	return -1
}
