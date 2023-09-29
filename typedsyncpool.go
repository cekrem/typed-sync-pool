package typedsyncpool

import (
	"sync"
)

// Pool is the same as a regular sync.Pool, but with generic typings
// (no need for manual type assertions)
type Pool[T any] struct {
	internalPool sync.Pool
}

// Put does the same as sync.Pool.Put, but infers type based on root pool type
func (t *Pool[T]) Put(x T) {
	t.internalPool.Put(x)
}

// Gut does the same as sync.Pool.Get, but infers type based on root pool type
func (t *Pool[T]) Get() T {
	return t.internalPool.Get().(T)
}

// New returns a new Pool; newFunc is technically optional, but highly recommended
//
// See sync.Pool docs for more
func New[T any](newFunc func() T) (tp Pool[T]) {
	if newFunc != nil {
		tp.internalPool.New = func() any {
			return newFunc()
		}
	}

	return
}
