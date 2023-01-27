package concurrentmap

import (
	"os"
	"reflect"
	"sync"
	"testing"
)

var tMap *CMap[int, struct{}]

const (
	firstKey, lastKey = 1, 100
	impossibleKey     = firstKey - 1
)

func TestMain(m *testing.M) {
	wg := sync.WaitGroup{}
	tMap = New[int, struct{}]()

	for i := firstKey; i <= lastKey; i++ {
		wg.Add(1)

		go func(key int) {
			tMap.Insert(key, struct{}{})
			wg.Done()
		}(i)
	}
	wg.Wait()

	code := m.Run()
	os.Exit(code)
}

func TestNew(t *testing.T) {
	t.Parallel()

	expMap := &CMap[int, struct{}]{
		mu: sync.RWMutex{},
		m:  make(map[int]struct{}),
	}

	resMap := New[int, struct{}]()

	if !reflect.DeepEqual(expMap, resMap) {
		t.Fatalf("Expected c-map: %v\nGot: %v", expMap, resMap)
	}
}

func TestCMapInsert(t *testing.T) {
	t.Parallel()

	for i := firstKey; i <= lastKey; i++ {
		_, ok := tMap.Get(i)
		if !ok {
			t.Fatalf("Expected non-empty entry "+
				"by key: %d\nGot: empty", i)
		}
	}
}

func TestCMapGet(t *testing.T) {
	t.Parallel()

	if value, ok := tMap.Get(impossibleKey); ok {
		t.Fatalf("Expected empty entry \nGot: %v, %v",
			impossibleKey, value)
	}

	e := make(chan int, lastKey)

	go func(err chan<- int) {
		for i := firstKey; i <= lastKey; i++ {
			go func(key int) {
				if _, ok := tMap.Get(key); !ok {
					err <- key
				} else {
					err <- impossibleKey
				}
			}(i)
		}
	}(e)

	for i := lastKey; i != 0; i-- {
		if key := <-e; key != impossibleKey {
			t.Fatalf("Expected non-empty entry "+
				"by key: %d\nGot: empty", key)
		}
	}
}
