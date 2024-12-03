package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolvePart1File(t *testing.T) {
	actual, err := SolvePart1File("test.txt")

	require.NoError(t, err)
	require.Equal(t, 11, actual)
}

func TestReadLists(t *testing.T) {
	actual, err := readLists("test.txt")

	expect := lists{
		left:  []int{3, 4, 2, 1, 3, 3},
		right: []int{4, 3, 5, 3, 9, 3},
	}
	require.NoError(t, err)
	require.Equal(t, expect, actual)
}

func TestSolvePart2File(t *testing.T) {
	actual, err := SolvePart2File("test.txt")

	require.NoError(t, err)
	require.Equal(t, 31, actual)
}
