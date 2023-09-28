package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunkAny(t *testing.T) {
	chunkSize := 3
	arrI64 := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arrI64Chunks := Chunk[int64](arrI64, chunkSize)
	t.Log(arrI64Chunks)
	assert.Equal(t, [][]int64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}}, arrI64Chunks)

	arrStr := []string{"a", "b", "c", "d", "e"}
	arrStrChunks := Chunk[string](arrStr, chunkSize)
	t.Log(arrStrChunks)
	assert.Equal(t, [][]string{{"a", "b", "c"}, {"d", "e"}}, arrStrChunks)
}

func TestMapGetOrDefault(t *testing.T) {
	m := map[string]int{
		"one": 1,
		"two": 2,
	}
	res1 := MapGetOrDefault[string, int](m, "one", 1)
	assert.Equal(t, 1, res1)

	res2 := MapGetOrDefault[string, int](m, "three", 3)
	assert.Equal(t, 3, res2)
}
