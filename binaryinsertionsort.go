package main

import "fmt"

func binarySearch(a []int, item int, low int, high int) int {
	if high <= low {
		if item > a[low] {
			return (low + 1)
		} else {
			return low
		}
	}

	mid := (low + high) / 2

	if item == a[mid] {
		return mid + 1
	}

	if item > a[mid] {
		return binarySearch(a, item, mid+1, high)
	}
	return binarySearch(a, item, low, mid-1)
}

// Function to sort an array a[] of size 'n'
func insertionSort(a []int, n int) {
	var loc, j, selected int

	for i := 1; i < n; i++ {
		j = i - 1
		selected = a[i]

		// find location where selected should be inseretd
		loc = binarySearch(a, selected, 0, j)

		// Move all elements after location to create space
		for j >= loc {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = selected
	}
}

// Driver Code
func main() {
	a := []int{37, 23, 0, 17, 12, 72, 31, 46, 100, 88, 54}
	n := len(a)

	insertionSort(a, n)

	fmt.Printf("Sorted array: \n")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a[i])
	}
}
