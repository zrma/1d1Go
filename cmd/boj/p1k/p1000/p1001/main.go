package main

import (
	"bufio"
	"os"

	"1d1go/boj/p1k/p1000"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	p1000.Solve1001(scanner, writer)
}
