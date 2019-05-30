package set

import (
	"log"
	"testing"
)

func TestStringSet(t *testing.T) {
	s := StringSetNew()
	v := "hello"
	s.Push(&v)
	v = "world"
	s.Push(&v)
	v = "hello"
	s.Push(&v)
	v = "world"
	s.Push(&v)
	l := s.List()
	if l != nil {
		log.Println(*l)
	}
}
