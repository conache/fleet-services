package repository

import (
	"fmt"
	"math"
)

type Repository interface{}

type CacheEntityType string

var maxChunkStoreSize uint64 = uint64(math.Pow(2, 32) - 1)

const (
	chunkEntity CacheEntityType = "chunk"
	fileEntity  CacheEntityType = "file"
	storeEntity CacheEntityType = "store"
)

func composeFileKey(fileId string) string {
	return fmt.Sprintf("%s:%s", fileEntity, fileId)
}

func composeChunkKey(chunkSha2 string) string {
	return fmt.Sprintf("%s:%s", chunkEntity, chunkSha2)
}

func composeFileChunkStoreKey(fileId string, chunkStoreIndex uint64) string {
	return fmt.Sprintf("%s:%d", composeFileKey(fileId), chunkStoreIndex)
}

func composeFileChunkBindingKey(fileId string, chunkSha2 string) string {
	return fmt.Sprintf("%s:%s", composeFileKey(fileId), composeChunkKey(chunkSha2))
}