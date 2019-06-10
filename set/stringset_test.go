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

func TestStringSetIsEqual(t *testing.T) {
	li1 := []string{"1", "2", "3"}
	li2 := []string{"2", "1", "3"}
	s1 := StringSetFromList(&li1)
	s2 := StringSetFromList(&li2)
	isEqual := s1.IsEqual(s2)
	log.Println(isEqual)
}
