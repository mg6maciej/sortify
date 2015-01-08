package sortify

import (
	"math/rand"
	"sort"
)

func Genetic(data sort.Interface) {
	currentEval := eval(data)
	for currentEval > 0 {
		mutations := mutations(data.Len())
		mutate(data, mutations)
		newEval := eval(data)
		if currentEval > newEval {
			currentEval = newEval
		} else {
			reverse(mutations)
			mutate(data, mutations)
		}
	}
}

func eval(data sort.Interface) int {
	eval := 0
	l := data.Len()
	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if data.Less(i, j) {
				eval++
			}
		}
	}
	return eval
}

func mutations(l int) []int {
	count := l/6 + 1
	mutations := make([]int, 2*count)
	for i := 0; i < 2*count; i++ {
		mutations[i] = rand.Intn(l)
	}
	return mutations
}

func mutate(data sort.Interface, mutations []int) {
	for i := 0; i < len(mutations); i += 2 {
		data.Swap(mutations[i], mutations[i+1])
	}
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
