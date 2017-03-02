package vpndns

import "sync"

type KeyValueStore interface {
    Get(key string) (string, bool)
    Put(key string, val string)
}

type StringMap map[string]string
type ConcurrentMap struct {
    m StringMap
    lock sync.Mutex
}

func NewConcurrentMap() ConcurrentMap {
    return ConcurrentMap { m: make(StringMap), lock: sync.Mutex{} }
}

func (c ConcurrentMap) Get(key string) (string, bool) {
    c.lock.Lock()
    defer c.lock.Unlock()
    val, ok := c.m[key]
    return val, ok
}

func (c ConcurrentMap) Put(key string, val string) {
    c.lock.Lock()
    defer c.lock.Unlock()
    c.m[key] = val
}
