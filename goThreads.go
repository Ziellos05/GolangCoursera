package main

import (
	"fmt"
	"sync"
)

func main() {
	// Sample unsorted array
	unsorted := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6, 1, 8, 4, 13}

	// Create a wait group to wait for the goroutines to finish
	var wg sync.WaitGroup

	// Number of goroutines
	numGoroutines := 4

	// Determine the size of each sub-array
	subArraySize := len(unsorted) / numGoroutines

	// Create a channel to receive sorted sub-arrays
	sortedCh := make(chan []int, numGoroutines)

	// Split the unsorted array into sub-arrays and sort them concurrently
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		start := i * subArraySize
		end := (i + 1) * subArraySize
		if i == numGoroutines-1 {
			end = len(unsorted)
		}
		go func(arr []int) {
			defer wg.Done()
			quickSort(arr)
			sortedCh <- arr
		}(unsorted[start:end])
	}

	// Close the channel after all goroutines are done
	go func() {
		wg.Wait()
		close(sortedCh)
	}()

	// Merge the sorted sub-arrays into a single sorted array
	sorted := mergeSortedArrays(sortedCh, numGoroutines)

	fmt.Println("Sorted Array:", sorted)
}

// quickSort sorts an array using the quicksort algorithm
func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivot := arr[len(arr)/2]
	left, right := 0, len(arr)-1

	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	quickSort(arr[:right+1])
	quickSort(arr[left:])
}

// mergeSortedArrays merges sorted sub-arrays into a single sorted array
func mergeSortedArrays(ch <-chan []int, num int) []int {
	result := make([]int, 0, num)

	for i := 0; i < num; i++ {
		subArray := <-ch
		result = merge(result, subArray)
	}

	return result
}

// merge merges two sorted arrays into a single sorted array
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
