package main

import (
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestUnique(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("should not have repeating values", prop.ForAll(
		func(islice []int) bool {
			unqSlice := unique(islice)
			for _, i := range unqSlice {
				// i = unqSlice.pop(), check rest of unqSlice for same number
				unqSlice = unqSlice[1:]
				for _, j := range unqSlice {
					if i == j {
						return false
					}
				}
			}
			return true
		},
		gen.SliceOf(gen.Int()),
	))
	properties.TestingRun(t)
}

func TestRemoveZeroes(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("should not have zeroes", prop.ForAll(
		func(islice []int) bool {
			nozSlice := removeZeroes(islice)
			for _, i := range nozSlice {
				if i == 0 {
					return false
				}
			}
			return true
		},
		gen.SliceOf(gen.Int()),
	))
	properties.TestingRun(t)
}
