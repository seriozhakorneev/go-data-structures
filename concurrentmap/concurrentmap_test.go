package concurrentmap

import (
	"os"
	"reflect"
	"sync"
	"testing"
)

const (
	firstKey, lastKey = 1, 100
	impossibleKey     = firstKey - 1

	tValue    = 0
	tUpdValue = tValue - 1
)

var tCMap *CMap[int, int]

type (
	CMapTest struct{ t *testing.T }

	writeResult struct {
		key   int
		exist bool
	}
)

func TestMain(m *testing.M) {
	wg := sync.WaitGroup{}
	tCMap = New[int, int]()

	for i := firstKey; i <= lastKey; i++ {
		wg.Add(1)
		go func(key int) {
			tCMap.Insert(key, tValue)
			wg.Done()
		}(i)
	}
	wg.Wait()

	code := m.Run()
	os.Exit(code)
}

func TestRun(t *testing.T) {
	t.Run("Insert/Read/Update", func(t *testing.T) {
		test := CMapTest{t}
		test.TestNew()
		test.TestInsert()
		test.TestGet()
		test.TestUpdate()
		test.TestGetMap()
	})

	t.Run("Delete", func(t *testing.T) {
		t.Parallel()
		test := CMapTest{t}
		test.TestDelete()
	})
}

func (c *CMapTest) TestNew() {
	expMap := &CMap[int, int]{
		mu: sync.RWMutex{},
		m:  make(map[int]int),
	}

	resMap := New[int, int]()

	if !reflect.DeepEqual(expMap, resMap) {
		c.t.Fatalf("Expected c-map: %v\nGot: %v", expMap, resMap)
	}
}

func (c *CMapTest) TestGetMap() {
	expType := reflect.TypeOf(map[int]int{})
	underMap := reflect.TypeOf(tCMap.getMap())

	if !underMap.AssignableTo(expType) {
		c.t.Fatalf("Expected type: %v\nGot: %v", expType, underMap)
	}
}

func (c *CMapTest) TestInsert() {
	for i := firstKey; i <= lastKey; i++ {
		_, ok := tCMap.Get(i)
		if !ok {
			c.t.Fatalf("Expected non-empty entry "+
				"by key: %d\nGot: empty", i)
		}
	}
}

func (c *CMapTest) TestGet() {
	if value, ok := tCMap.Get(impossibleKey); ok {
		c.t.Fatalf("Expected empty entry \nGot: %v, %v", impossibleKey, value)
	}

	e := make(chan int, lastKey)

	go func(err chan<- int) {
		for i := firstKey; i <= lastKey; i++ {
			go func(key int) {
				if _, ok := tCMap.Get(key); !ok {
					err <- key
				} else {
					err <- impossibleKey
				}
			}(i)
		}
	}(e)

	for i := lastKey; i != 0; i-- {
		if key := <-e; key != impossibleKey {
			c.t.Fatalf("Expected non-empty entry by key: %d\nGot: empty", key)
		}
	}
}

func (c *CMapTest) TestUpdate() {
	if ok := tCMap.Update(impossibleKey, tUpdValue); ok {
		c.t.Fatalf(
			"Expected update result by key(%d): false\nGot: %t",
			impossibleKey,
			ok,
		)
	}

	e := make(chan *writeResult, lastKey)

	go func(err chan<- *writeResult) {
		for i := firstKey; i <= lastKey; i++ {
			go func(key int) {
				if ok := tCMap.Update(key, tUpdValue); !ok {
					err <- &writeResult{key, ok}
				} else {
					err <- nil
				}
			}(i)
		}
	}(e)

	for i := firstKey; i <= lastKey; i++ {
		if result := <-e; result != nil {
			c.t.Fatalf("Failed to update entry by key: %d\nGot: %t", result.key, result.exist)
		}
	}

	for i := firstKey; i <= lastKey; i++ {
		val, ok := tCMap.Get(i)
		if !ok {
			c.t.Fatalf("Failed to update entry by key: %d\nGot: empty", i)
		}
		if val != tUpdValue {
			c.t.Fatalf("Failed to update entry by %d: %v\nGot: %d: %v", i, tUpdValue, i, val)
		}
	}
}

func (c *CMapTest) TestDelete() {
	if ok := tCMap.Delete(impossibleKey); ok {
		c.t.Fatalf("Expected delete result by key(%d): false\nGot: %t", impossibleKey, ok)
	}

	e := make(chan *writeResult, lastKey)
	go func(err chan<- *writeResult) {
		for i := firstKey; i <= lastKey; i++ {
			go func(key int) {
				if ok := tCMap.Delete(key); !ok {
					err <- &writeResult{key, ok}
				} else {
					err <- nil
				}
			}(i)
		}
	}(e)

	for i := firstKey; i <= lastKey; i++ {
		if result := <-e; result != nil {
			c.t.Fatalf("Failed to delete entry by key: %d\nGot: %t", result.key, result.exist)
		}
	}

	if l := len(tCMap.getMap()); l > 0 {
		c.t.Fatalf("Expected map len: %d\nGot: %d", 0, l)
	}
}
