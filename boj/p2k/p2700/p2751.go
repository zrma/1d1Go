package p2700

import (
	"fmt"
	"strconv"
)

func Solve2751(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n)
	for i := range arr {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	mergeSort(arr)

	for i := range arr {
		_, _ = fmt.Fprintln(writer, arr[i])
	}
}

func mergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2
	left := make([]int, mid)
	right := make([]int, len(arr)-mid)

	copy(left, arr[:mid])
	copy(right, arr[mid:])

	mergeSort(left)
	mergeSort(right)

	merge(arr, left, right)
}

func merge(arr []int, left []int, right []int) {
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}