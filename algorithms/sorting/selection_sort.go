package main

/*
 * https://en.wikipedia.org/wiki/Selection_sort
 * Time Complexity: O(n*n) as there are two nested loops.
 * Auxiliary Space: O(1)
 */

import "fmt"

func main() {
  arr := [7] int{13, 2, 11, 7, 4, 12, 17}
  fmt.Println("Initial array is:", arr)

  l := len(arr)
  for i := 0; i < l - 1; i++ {

	/* Find the minimum element in unsorted array */
	min_idx := i
	for j := i + 1; j < l; j++ {
	  if arr[j] < arr[min_idx] {
        min_idx = j
	  }
	}

	/* Swap the found minimum element with the first element */
	swap(&arr[min_idx], &arr[i]);
  }

  fmt.Println("Sorted array is: ", arr)
}

func swap (x,y *int) {
    *x, *y = *y, *x
}
