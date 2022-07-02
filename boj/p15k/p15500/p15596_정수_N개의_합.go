package p15500

import (
	"fmt"
)

func Solve15596(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int, n)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	res := sum(arr)
	_, _ = fmt.Fprint(writer, res)
}

func sum(arr []int) int {
	var res int
	for _, n := range arr {
		res += n
	}
	return res
}
