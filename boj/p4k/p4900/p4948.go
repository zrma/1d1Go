package p4900

import (
	"fmt"
	"strconv"
)

func Solve4948(scanner Scanner, writer Writer) {
	sieveOfEratosthenes := buildSieveOfEratosthenes()
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			return
		}
		cnt := 0
		for i := n + 1; i <= 2*n; i++ {
			if !sieveOfEratosthenes[i] {
				cnt++
			}
		}
		_, _ = fmt.Fprintln(writer, cnt)
	}
}

const maxLen = 123456*2 + 1

func buildSieveOfEratosthenes() [maxLen]bool {
	sieveOfEratosthenes := [maxLen]bool{}
	for i := 2; i*i < maxLen; i++ {
		if !sieveOfEratosthenes[i] {
			for j := i * i; j < maxLen; j += i {
				sieveOfEratosthenes[j] = true
			}
		}
	}
	sieveOfEratosthenes[0] = true
	sieveOfEratosthenes[1] = true
	return sieveOfEratosthenes
}
