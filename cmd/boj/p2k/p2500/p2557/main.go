package main

import (
	"bufio"
	"os"

	"1d1go/boj/p2k/p2500"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	p2500.Solve2557(writer)
}
