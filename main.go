package main

import (
	"imgoptimizer/convert"
	"imgoptimizer/database"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args[1:]
	src := args[0]
	dst := args[1]
	convert.ToWebp(src, dst)
	_, fileName := filepath.Split(src)
	database.Add(fileName)
}
