package main

import (
  "fmt"
)

/*
 * Quick sort - https://en.wikipedia.org/wiki/Quicksort
*/

func main() {
  arr:= []int{13, 2, 11, 7, 4, 12, 17}
  fmt.Println("Initial array is:", arr)

  length := len(arr)
  fmt.Println("Sorted array is: ", quickSort(arr, 0, length - 1))
}

func quickSort(arr []int, l, h int) []int {

  if len(arr) <= 1 {
	return arr
  }

  if l < h {
	/* pi is partitioning index, arr[p] is now
	    at right place */
	pi := partition(arr, l, h)

	// Separately sort elements before
	// partition and after partition
	quickSort(arr, l, pi - 1)
	quickSort(arr, pi + 1, h)
  }
  return arr
}

/* This function takes last element as pivot, places
   the pivot element at its correct position in sorted
   array, and places all smaller (smaller than pivot)
   to left of pivot and all greater elements to right
   of pivot 
*/
func partition (arr []int, low, high int) int {

  pivot := arr[high]    // pivot
  i := (low - 1)  // Index of smaller element

  for j := low; j<= high - 1; j++ {
	// If current element is smaller than or
	// equal to pivot
	if arr[j] <= pivot {
	  i++ // increment index of smaller element
	  swap(&arr[i], &arr[j])
	}
  }

  swap(&arr[i + 1], &arr[high])
  return (i + 1)
}

// A utility function to swap two elements
func swap (x,y *int) {
  *x, *y = *y, *x
}
