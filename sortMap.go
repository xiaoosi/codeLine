package main

type item struct {
	key   string
	value int
}
type sortMap []item

func convTo(from map[string]int) sortMap {
	res := sortMap{}
	for k, v := range from {
		res = append(res, item{
			key:   k,
			value: v,
		})
	}
	return res
}

func (s sortMap) Len() int {
	return len(s)
}

func (s sortMap) Less(i, j int) bool {
	if s[i].value < s[j].value {
		return true
	} else {
		return false
	}
}

func (s sortMap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

