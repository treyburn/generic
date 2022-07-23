package cache

import (
	"github.com/stretchr/testify/assert"
	"sync"
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

// be sure to add --race flag for testing
func Test_LRU_Concurrency(t *testing.T) {
	var wg sync.WaitGroup
	lru := NewLRU[int, int](10)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(l *LRUCache[int, int]) {
			for j := 0; j < 10000; j++ {
				switch j % 2 {
				case 1:
					l.Put(j, j)
				default:
					var err error
					_, err = l.Get(j - 1)
					if err != nil {
						continue
					}
				}
			}
			wg.Done()
		}(lru)
	}
	wg.Wait()
}
