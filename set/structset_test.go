package set

import (
    "testing"
    "log"
)

type config struct {
    projectUuid string
    serviceName string
}

func TestStructSet(t *testing.T) {
    s := StructSetNew()
    c := config{
        projectUuid: "123",
        serviceName: "456",
    }
    s.Push(c)
    s.Push(c)
    if s.Exists(c) {
        log.Println("exist")
    } else {
        log.Println("not exist")
    }
    s.Range(func(v interface{}) bool {
        log.Println(v.(config))
        return true
    })
    log.Println(s.Len())
}
