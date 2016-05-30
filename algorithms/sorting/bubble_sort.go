package main

/*
 * Bubble sort - http://en.wikipedia.org/wiki/Bubble_sort
 * Worst and Average Case Time Complexity: O(n*n). Worst case occurs when array is reverse sorted.
 * Best Case Time Complexity: O(n). Best case occurs when array is already sorted.
 * Auxiliary Space: O(1)
 */

import "fmt"

func main() {
  arr := [7]int{13, 2, 11, 7, 4, 12, 17}
  fmt.Println("Initial array is:", arr)

  l := len(arr)
  for i := 0; i < l - 1; i++ {
	for j := 0; j < l - i - 1; j++ {
	  if arr[j] > arr[j + 1] {
		swap(&arr[j], &arr[j + 1])
	  }
	}
  }

  fmt.Println("Sorted array is: ", arr)
}

func swap (x,y *int) {
  *x, *y = *y, *x
}
