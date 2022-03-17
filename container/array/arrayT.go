package array

import (
	"fmt"
	"sync"
)

type ArrayT[T comparable] struct {
	Slice []T
	lock  sync.RWMutex // 加锁
}

func NewT[T comparable]() *ArrayT[T] {
	return &ArrayT[T]{Slice: make([]T, 0)}
}

func (a *ArrayT[T]) Append(value T) {
	a.lock.Lock()
	defer a.lock.Unlock()

	a.Slice = append([]T(a.Slice), value)
}

func (a *ArrayT[T]) Len() int {
	a.lock.RLock()
	defer a.lock.RUnlock()

	return len([]T(a.Slice))
}

func (a *ArrayT[T]) List() []T {
	a.lock.RLock()
	defer a.lock.RUnlock()

	return a.Slice
}

//Insert Array
func (a *ArrayT[T]) Insert(index int, value T) {
	a.lock.Lock()
	defer a.lock.Unlock()

	var reset = make([]T, 0)
	prefix := append(reset, a.Slice[index:]...)
	a.Slice = append([]T(a.Slice[0:index]), value)
	a.Slice = append([]T(a.Slice), prefix...)
}

func (a *ArrayT[T]) Delete(index int) (t T, err error) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if index < 0 || index >= len([]T(a.Slice)) {
		err = fmt.Errorf("invalid index %d", index)
		return
	}
	t = a.Slice[index]
	a.Slice = append([]T(a.Slice[:index]), a.Slice[index+1:]...)
	return
}

func (a *ArrayT[T]) Set(index int, value T) bool {
	a.lock.Lock()
	defer a.lock.Unlock()

	if index < 0 || index >= len([]T(a.Slice)) {
		return false
	}
	a.Slice[index] = value
	return true
}

func (a *ArrayT[T]) Get(index int) (t T, err error) {
	a.lock.RLock()
	defer a.lock.RUnlock()

	if index < 0 || index >= len([]T(a.Slice)) {
		err = fmt.Errorf("invalid index %d", index)
		return
	}
	return a.Slice[index], nil
}

func (a *ArrayT[T]) Search(value T) int {
	a.lock.RLock()
	defer a.lock.RUnlock()

	for i, v := range a.Slice {
		if v == value {
			return i
		}
	}
	return -1
}

func (a *ArrayT[T]) Clear() {
	a.lock.Lock()
	defer a.lock.Unlock()

	a.Slice = make([]T, 0)
}

func (a *ArrayT[T]) LockFunc(f func(array []T)) *ArrayT[T] {
	a.lock.Lock()
	defer a.lock.Unlock()

	f(a.Slice)
	return a
}
