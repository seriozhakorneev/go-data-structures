package concurrentmap

import (
	"sync"
)

/*
Реализовать библиотеку универсального key-value хранилища.

Основные требования к реализации:

Должны поддерживаться все сравниваемые типы данных, т.е. хранилище должно уметь
работать с ключами и значениями типа interface{}.
Должны поддерживаться следующие операции: вставка, получение, обновление и удаление.
Хранилище должно быть потокобезопасным (согласованное чтение и запись из разных горутин).
*/

// CMap - concurrent-safety map
type CMap[K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
	// generic zero value
	zero V
}

// New - returns new CMap exemplar.
func New[K comparable, V any]() *CMap[K, V] {
	return &CMap[K, V]{
		mu: sync.RWMutex{},
		m:  make(map[K]V),
	}
}

// Insert -
func (c *CMap[K, V]) Insert(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.m[key] = value
}

func (c *CMap[K, V]) Get(key K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if value, ok := c.m[key]; ok {
		return value, true
	}

	return c.zero, false
}

//
//func (m *MyStorage) Update(k, value any) (bool, error) {
//	m.mu.Lock()
//	defer m.mu.Unlock()
//
//	if _, ok := m.m[k]; !ok {
//		return false, KeyNotFound
//	}
//
//	m.m[k] = value
//	return true, nil
//}
//
//func (m *MyStorage) Delete(k any) error {
//	m.mu.Lock()
//	defer m.mu.Unlock()
//	delete(m.m, k)
//
//	return nil
//}
