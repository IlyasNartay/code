package main

import (
	"fmt"
	"sort"
)

func main() {
	a := [7]int{6, 3, 5, 6, 3, 5, 7}
	b := a[2:7]
	SortSlice(&b)
	fmt.Print(b)
}

func SortSlice(slice *[]int) {
	sort.Slice(*slice, func(i, j int) bool { return i < j })
}
