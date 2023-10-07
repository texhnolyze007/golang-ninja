package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func AddOrRemove(aPtr *[]int, elem int) {
	for i := 0; i < len(*aPtr)-1; i++ {
		if (*aPtr)[i] == elem {
			(*aPtr)[i], (*aPtr)[i+1] = (*aPtr)[i+1], (*aPtr)[i]
		}
	}

	if (*aPtr)[len(*aPtr)-1] == elem {
		*aPtr = (*aPtr)[:len(*aPtr)-1]
	} else {
		*aPtr = append(*aPtr, elem)
	}
}

func TestAddOrRemove(t *testing.T) {
	a := []int{1, 2, 3, 4}

	for i := 0; i < 20; i++ {
		AddOrRemove(&a, i)
	}

	expected := []int{0, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	require.Equal(t, expected, a)
}
