package main

import (
	"bufio"
	"flag"
	"os"
	"path/filepath"
)

var listWriter *bufio.Writer

func WalkFuncByFile(path string, info os.FileInfo, err error) error {
	var writer_err error

	if info.IsDir() == false {
		listWriter.WriteString(path)
		listWriter.WriteString("\r\n")
		writer_err = listWriter.Flush()
	}

	return writer_err
}

func WalkFuncByDirectory(path string, info os.FileInfo, err error) error {
	var writer_err error

	if info.IsDir() == true {
		listWriter.WriteString(path)
		listWriter.WriteString("\r\n")
		writer_err = listWriter.Flush()
	}

	return writer_err
}

func main() {
	var path = flag.String("path", "./", "directory path")
	var searchTarget = flag.String("target", "file", "file or directory")

	flag.Parse()
	resultFile := "./result.txt"
	f, err := os.Create(resultFile)
	if err != nil {
		panic(err)
	}

	listWriter = bufio.NewWriter(f)

	if *searchTarget == "file" {
		filepath.Walk(*path, WalkFuncByFile)
	} else if *searchTarget == "directory" {
		filepath.Walk(*path, WalkFuncByDirectory)
	}
}
