package set

type CStringSet struct {
	m map[string]int
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

func StringSetNew() *CStringSet {
	s := CStringSet{
		m: make(map[string]int),
	}
	return &s
}
