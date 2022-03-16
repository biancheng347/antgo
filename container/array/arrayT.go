package array

import "sync"

type ArrayT[T any] struct {
	Slice []T
	lock  sync.RWMutex // 加锁
}

func NewT[T any]() *ArrayT[T] {
	return &ArrayT[T]{Slice: make([]T, 0)}
}

func (a *ArrayT[T]) Append(value T) {
	a.lock.Lock()
	defer a.lock.Unlock()

	a.Slice = append(a.Slice, value)
}

/*

//Len Count Array
func (a *Array) Len() int {
	a.lock.RLock()
	defer a.lock.RUnlock()

	return len(a.Slice)
}

//List Array
func (a *Array) List() []interface{} {
	a.lock.RLock()
	defer a.lock.RUnlock()

	return a.Slice
}

//Insert Array
func (a *Array) Insert(index int, value interface{}) {
	a.lock.Lock()
	defer a.lock.Unlock()

	var reset = make([]interface{}, 0)
	prefix := append(reset, a.Slice[index:]...)
	a.Slice = append(a.Slice[0:index], value)
	a.Slice = append(a.Slice, prefix...)
}

//Delete Array
func (a *Array) Delete(index int) interface{} {
	a.lock.Lock()
	defer a.lock.Unlock()

	if index < 0 || index >= len(a.Slice) {
		return nil
	}
	value := a.Slice[index]
	a.Slice = append(a.Slice[:index], a.Slice[index+1:]...)
	return value
}

//Set Array
func (a *Array) Set(index int, value interface{}) bool {
	a.lock.Lock()
	defer a.lock.Unlock()

	if index < 0 || index >= len(a.Slice) {
		return false
	}
	a.Slice[index] = value
	return true
}

//Get Array
func (a *Array) Get(index int) interface{} {
	a.lock.RLock()
	defer a.lock.RUnlock()

	if index < 0 || index >= len(a.Slice) {
		return nil
	}
	return a.Slice[index]
}

//Search Array
func (a *Array) Search(value interface{}) int {
	a.lock.RLock()
	defer a.lock.RUnlock()

	for i, v := range a.Slice {
		if v == value {
			return i
		}
	}
	return -1
}

//Clear Array
func (a *Array) Clear() {
	a.lock.Lock()
	defer a.lock.Unlock()

	a.Slice = make([]interface{}, 0)
}

//LockFunc locks writing by callback function <f>
func (a *Array) LockFunc(f func(array []interface{})) *Array {
	a.lock.Lock()
	defer a.lock.Unlock()

	f(a.Slice)
	return a
}

*/

//ReadLockFunc locks writing by callback function <f>
/*
func (a *Array) ReadLockFunc(f func(array []interface{})) *Array {
	a.lock.RLock()
	defer a.lock.RUnlock()

	f(a.Slice)
	return a
}
*/
