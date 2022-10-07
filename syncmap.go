package syncmap

import "sync"

// sync.Map的泛型封装
type SyncMap[KeyT, ValueT any] struct {
	syncMap sync.Map
}

func (p SyncMap[KeyT, ValueT]) Load(key KeyT) (ValueT, bool) {
	value, ok := p.syncMap.Load(key)
	return value.(ValueT), ok
}
func (p SyncMap[KeyT, ValueT]) Store(key KeyT, value ValueT) {
	p.syncMap.Store(key, value)
}
func (p SyncMap[KeyT, ValueT]) LoadOrStore(key KeyT, value ValueT) (ValueT, bool) {
	actual, loaded := p.syncMap.LoadOrStore(key, value)
	return actual.(ValueT), loaded
}

func (p SyncMap[KeyT, ValueT]) LoadAndDelete(key KeyT) (ValueT, bool) {
	value, ok := p.syncMap.LoadAndDelete(key)
	return value.(ValueT), ok
}

func (p SyncMap[KeyT, ValueT]) Delete(key KeyT) {
	p.syncMap.Delete(key)
}
func (p SyncMap[KeyT, ValueT]) Range(f func(key KeyT, value ValueT) bool) {
	p.syncMap.Range(func(key, value any) bool {
		return f(key.(KeyT), value.(ValueT))
	})
}
