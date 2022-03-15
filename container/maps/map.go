package maps

import (
	"sync"
)

type Key = interface{}

//Map parameter structure
type Map struct {
	Map  map[Key]interface{} //
	lock sync.RWMutex        // 加锁
}

//New ...
func New() *Map {
	return &Map{Map: make(map[Key]interface{})}
}

//Set ...
func (m *Map) Set(key Key, value interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.Map[key] = value
}

//Get ...
func (m *Map) Get(key Key) interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()

	v, _ := m.Map[key]
	return v
}

//GetOrSet ...
func (m *Map) GetOrSet(key Key, value interface{}) interface{} {
	m.lock.Lock()
	defer m.lock.Unlock()

	v, ok := m.Map[key]
	if !ok {
		m.Map[key] = value
		return value
	}
	return v
}

//Count ...
func (m *Map) Count() int {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return len(m.Map)
}

//Delete ...
func (m *Map) Delete(key Key) bool {
	m.lock.Lock()
	defer m.lock.Unlock()

	delete(m.Map, key)
	_, ok := m.Map[key]
	return !ok
}

//LockFunc locks writing by callback function <f>
func (m *Map) LockFunc(f func(Map map[Key]interface{})) *Map {
	m.lock.Lock()
	defer m.lock.Unlock()

	f(m.Map)
	return m
}

/*
//ReadLockFunc locks writing by callback function <f>
func (m *Map) ReadLockFunc(f func(Map map[interface{}]interface{})) *Map {
	m.lock.RLock()
	defer m.lock.RUnlock()

	f(m.Map)
	return m
}
*/
