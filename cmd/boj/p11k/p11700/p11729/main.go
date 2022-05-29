package main

import (
	"bufio"
	"os"

	"1d1go/boj/p11k/p11700"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	p11700.Solve11729(scanner, writer)
}
