package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunkAny(t *testing.T) {
	chunkSize := 3
	arrI64 := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arrI64Chunks := ChunkAny[int64](arrI64, chunkSize)
	t.Log(arrI64Chunks)
	assert.Equal(t, [][]int64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}}, arrI64Chunks)

	arrStr := []string{"a", "b", "c", "d", "e"}
	arrStrChunks := ChunkAny[string](arrStr, chunkSize)
	t.Log(arrStrChunks)
	assert.Equal(t, [][]string{{"a", "b", "c"}, {"d", "e"}}, arrStrChunks)
}
