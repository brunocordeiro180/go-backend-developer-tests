package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("SP// Backend Developer Test - Input Processing")
	fmt.Println()

	// Read STDIN into a new buffered reader
	reader := bufio.NewScanner(os.Stdin)

	for reader.Scan() {
		line := reader.Text()
		if strings.Contains(strings.ToLower(line), "error") {
			fmt.Println(line)
		}

		if err := reader.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading standard input:", err)
		}
	}
}
