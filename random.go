package sortify

import (
	"math/rand"
	"sort"
)

func Random(data sort.Interface) {
	l := data.Len()
	for !sorted(data) {
		i := rand.Intn(l)
		j := rand.Intn(l)
		data.Swap(i, j)
	}
}

func sorted(data sort.Interface) bool {
	l := data.Len()
	for i := 1; i < l; i++ {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}
