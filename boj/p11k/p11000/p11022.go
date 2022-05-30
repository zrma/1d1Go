package p11000

import (
	"bufio"
	"fmt"
	"os"
)

func Solve11022(arr2D [][2]int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() { _ = w.Flush() }()

	for i, arr := range arr2D {
		_, _ = fmt.Fprintf(w, "Case #%d: %d + %d = %d\n", i+1, arr[0], arr[1], arr[0]+arr[1])
	}
}