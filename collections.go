package main

// From https://www.golangprograms.com/remove-duplicate-values-from-slice.html
// Takes a slice and returns only unique values
func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func removeZeroes(intSlice []int) []int {
	r := []int{}
	for i := range intSlice {
		if intSlice[i] != 0 {
			r = append(r, intSlice[i])
		}
	}
	return r
}
