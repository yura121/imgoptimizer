package main

import (
	"imgoptimizer/convert"
	"os"
)

func main() {
	args := os.Args[1:]
	convert.ToWebp(args[0], args[1])
}
