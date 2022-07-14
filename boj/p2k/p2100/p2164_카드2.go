package p2100

import (
	"fmt"
)

func Solve2164(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int, n)
	for i := range arr {
		arr[i] = i + 1
	}

	for len(arr) > 1 {
		arr = arr[1:]
		arr = append(arr, arr[0])
		arr = arr[1:]
	}
	_, _ = fmt.Fprint(writer, arr[0])
}
