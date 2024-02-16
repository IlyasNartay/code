package main

// 220107151
// Iliyas Nartay
import (
	"fmt"
	"sort"
)

func main() {
	a := [7]int{6, 3, 5, 6, 3, 5, 7}
	b := a[2:7]
	fmt.Println("Our slice:", b)
	SortSlice(b)
	fmt.Println("Sorted Slice:", b)
	IncrementOdd(b)
	fmt.Println("Increased in odd position Slice:", b)
	PrintSlice(b)
	RevereSlice(b)
	fmt.Print("Reversed Slice: ", b)
}

func SortSlice(slice []int) {
	sort.Ints(slice)
}

func IncrementOdd(slice []int) {
	for i := 1; i < len(slice); i += 2 {
		slice[i]++
	}
}
func PrintSlice(slice []int) {
	fmt.Print("We can print by func: {")
	for _, element := range slice {
		fmt.Printf("%d ", element)
	}
	fmt.Println("}")
}
func RevereSlice(slice []int) {
	i := 0
	j := len(slice) - 1
	for i < j {
		temp := slice[i]
		slice[i] = slice[j]
		slice[j] = temp
		i++
		j--
	}
}
