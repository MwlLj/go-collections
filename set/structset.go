package set

import (
    "sync"
)

type StructSet struct {
    sets sync.Map
    length int64
    mutex sync.Mutex
}

func (self *StructSet) Push(v interface{}) {
    _, b := self.sets.LoadOrStore(v, 0)
    if !b {
        self.mutex.Lock()
        self.length += 1
        self.mutex.Unlock()
    }
}

func (self *StructSet) Remove(v interface{}) {
    if self.Exists(v) {
        self.sets.Delete(v)
        self.mutex.Lock()
        self.length -= 1
        self.mutex.Unlock()
    }
}

func (self *StructSet) Exists(v interface{}) bool {
    _, b := self.sets.Load(v)
    return b
}

func (self *StructSet) Range(f func(v interface{}) bool) {
    self.sets.Range(func(k interface{}, v interface{}) bool {
        return f(k)
    })
}

func (self *StructSet) Len() int64 {
    return self.length
}

func StructSetNew() *StructSet {
    return &StructSet{
        length: 0,
    }
}
