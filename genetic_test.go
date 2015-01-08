package sortify

import (
	"fmt"
	"sort"
)

func ExampleGenetic() {
	a := []int{3, 21, 8, 7, 32, 1, 4, 3, 21}
	Genetic(sort.IntSlice(a))
	fmt.Print(a)
	//output: [1 3 3 4 7 8 21 21 32]
}
