package main

import (
	"fmt"
	"math"
)

func MainTruc() {
	fmt.Println("Number to truncate")
	var number float64
	// var truncateNumber string
	fmt.Scan(&number)
	integer := int(math.Trunc(number))
	fmt.Printf("%d", integer)
}
