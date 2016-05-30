package main

/*
 * Time Complexity: O(nLogn) in all 3 cases (worst, average and best) as merge sort always divides the array in two halves and take linear time to merge two halves.
 * Auxiliary Space: O(n)
 * Algorithmic Paradigm: Divide and Conquer
 * Sorting In Place: No in a typical implementation
 * Stable: Yes
 */

import "fmt"

func main() {
	arr := []int{13, 2, 11, 7, 4, 12, 17}
	fmt.Println("Initial array is:", arr)

	arr_size := len(arr)
	mergeSort(arr, 0, arr_size-1)
	fmt.Println("Sorted array is: ", arr)
}

/* l is for left index and r is right index of the
   sub-array of arr to be sorted */
func mergeSort(arr []int, l, r int) {

	if l < r {

		// Same as (l+r)/2, but avoids overflow for large l and r
		m := l + (r-l)/2

		// Sort first and second halves
		mergeSort(arr, l, m)
		mergeSort(arr, m+1, r)

		merge(arr, l, m, r)
	}
}

func merge(a []int, l, m, r int) {

	n1 := m - l + 1
	n2 := r - m

	// create temp arrays
	L := make([]int, n1)
	R := make([]int, n2)

	// Copy data to temp arrays L[] and R[]
	for i := 0; i < n1; i++ {
		L[i] = a[l+i]
	}

	for j := 0; j < n2; j++ {
		R[j] = a[m+1+j]
	}

	// Merge the temp arrays back into arr[l..r]
	i := 0 // Initial index of first subarray
	j := 0 // Initial index of second subarray
	k := l // Initial index of merged subarray

	for i < n1 && j < n2 {

		if L[i] <= R[j] {
			a[k] = L[i]
			i++
		} else {
			a[k] = R[j]
			j++
		}
		k++
	}

	// Copy the remaining elements of L[], if there are any
	for i < n1 {
		a[k] = L[i]
		i++
		k++
	}

	for j < n2 {
		a[k] = R[j]
		j++
		k++
	}

}
