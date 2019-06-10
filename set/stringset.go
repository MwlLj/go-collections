package set

import (
	"bytes"
)

type CStringSet struct {
	m map[string]int
}

func (this *CStringSet) IsEqual(other *CStringSet) bool {
	if len(this.m) != len(other.m) {
		return false
	}
	ret := true
	for key, _ := range this.m {
		if _, ok := other.m[key]; !ok {
			ret = false
			break
		}
	}
	return ret
}

func (this *CStringSet) Push(s *string) {
	if s == nil {
		return
	}
	if _, ok := this.m[*s]; ok {
		this.m[*s] += 1
	} else {
		this.m[*s] = 0
	}
}

func (this *CStringSet) List() *[]string {
	l := []string{}
	for key, _ := range this.m {
		l = append(l, key)
	}
	return &l
}

func (this CStringSet) String() string {
	buffer := bytes.Buffer{}
	for key, _ := range this.m {
		buffer.WriteString(key)
	}
	return buffer.String()
}

func StringSetNew() *CStringSet {
	s := CStringSet{
		m: make(map[string]int),
	}
	return &s
}

func StringSetFromList(li *[]string) *CStringSet {
	if li == nil {
		return nil
	}
	s := CStringSet{
		m: make(map[string]int),
	}
	for _, item := range *li {
		(s.m)[item] = 0
	}
	return &s
}
