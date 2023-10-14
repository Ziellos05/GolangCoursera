package main

import "fmt"

func MainBubbleSort() {
	var slice []int

	fmt.Println(slice)

	for i := 0; i < 10; i++ {
		var number int
		fmt.Scan(&number)
		slice = append(slice, number)
	}

	BubbleSort(slice...)

}

func BubbleSort(slice ...int) {
	var aux int
	for i := 1; i < len(slice); i++ {
		for j := 0; j < len(slice)-i; j++ {
			if slice[j] > slice[j+1] {
				aux = slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = aux
			}
		}
	}
	fmt.Println(slice)
}
