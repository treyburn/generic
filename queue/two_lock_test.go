package queue

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestTwoLockFIFO_Enqueue(t *testing.T) {
	want1 := "foo"
	want2 := "bar"
	q := NewFIFO[string]()
	assert.Equal(t, q.head, q.tail)
	assert.Nil(t, q.tail.Next)

	q.Enqueue(want1)
	assert.Equal(t, want1, q.head.Next.Value)
	assert.Equal(t, want1, q.tail.Value)
	assert.Nil(t, q.tail.Next)

	q.Enqueue(want2)
	assert.Equal(t, want1, q.head.Next.Value)
	assert.Equal(t, want2, q.head.Next.Next.Value)
	assert.Equal(t, want2, q.tail.Value)
	assert.Equal(t, q.tail, q.head.Next.Next)
	assert.Nil(t, q.tail.Next)
}

func TestTwoLockFIFO_Dequeue(t *testing.T) {
	want1 := "foo"
	want2 := "bar"

	q := NewFIFO[string]()

	got := q.Dequeue()
	assert.Empty(t, got)

	q.Enqueue(want1)
	got = q.Dequeue()
	assert.Equal(t, want1, got)

	got = q.Dequeue()
	assert.Empty(t, got)
	got = q.Dequeue()
	assert.Empty(t, got)

	q.Enqueue(want1)
	q.Enqueue(want2)

	got = q.Dequeue()
	assert.Equal(t, want1, got)
	got = q.Dequeue()
	assert.Equal(t, want2, got)

	got = q.Dequeue()
	assert.Empty(t, got)
}

func TestTwoLockFIFO_Concurrency(t *testing.T) {
	var wg sync.WaitGroup
	q := NewFIFO[string]()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			var got = "buzz"
			for j := 0; j < 10000; j++ {
				switch j % 3 {
				case 2:
					got = q.Dequeue()
				case 1:
					got = "fizz"
					q.Enqueue(got)
				default:
					q.Enqueue(got)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
