package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LRUCache(t *testing.T) {
	lru := NewLRU[int, int](2)
	lru.Put(1, 1)
	assert.Equal(t, 1, len(lru.cache))

	lru.Put(2, 2)
	assert.Equal(t, 2, len(lru.cache))

	got, err := lru.Get(1)
	assert.Equal(t, 1, got)
	assert.NoError(t, err)

	lru.Put(3, 3)
	assert.Equal(t, 2, len(lru.cache))

	got, err = lru.Get(2)
	assert.Equal(t, 0, got)
	assert.ErrorIs(t, err, ErrNotFound)

	lru.Put(4, 4)
	assert.Equal(t, 2, len(lru.cache))

	got, err = lru.Get(1)
	assert.Equal(t, 0, got)
	assert.ErrorIs(t, err, ErrNotFound)

	got, err = lru.Get(3)
	assert.Equal(t, 3, got)
	assert.NoError(t, err)

	got, err = lru.Get(4)
	assert.Equal(t, 4, got)
	assert.NoError(t, err)

	lru.Put(3, 4)

	lru.Put(5, 1)

	got, err = lru.Get(3)
	assert.Equal(t, 4, got)
	assert.NoError(t, err)

	got, err = lru.Get(3)
	assert.Equal(t, 4, got)
	assert.NoError(t, err)

	got, err = lru.Get(3)
	assert.Equal(t, 4, got)
	assert.NoError(t, err)

	got, err = lru.Get(3)
	assert.Equal(t, 4, got)
	assert.NoError(t, err)

	lru.Put(3, 5)

	got, err = lru.Get(3)
	assert.Equal(t, 5, got)
	assert.NoError(t, err)

	got, err = lru.Get(5)
	assert.Equal(t, 1, got)
	assert.NoError(t, err)
}
