package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/lucasjones/reggen"
)

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		println(err.Error())
	}

	if (fi.Mode() & os.ModeCharDevice) != 0 {
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str, err := reggen.Generate(scanner.Text(), 1)
		if err != nil {
			println(err.Error())
		}

		fmt.Println(str)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
