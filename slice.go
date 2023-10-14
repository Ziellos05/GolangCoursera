package main

import (
	"fmt"
	"sort"
)

func MainSlice() {

	slice := make([]int, 3)
	slice = slice[:0]

	var integer int

	for {
		n, err := fmt.Scan(&integer)
		if n == 1 || err == nil {
			slice = append(slice, integer)
			sort.Sort(sort.IntSlice(slice))
			printSlice(slice)
		} else {
			break
		}
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
