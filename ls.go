package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func isDotName(file os.FileInfo) bool {
	fileNameRune := []rune(file.Name())
	return fileNameRune[0] == rune('.')
}

func main() {
	var dirs []string

	if len(os.Args) == 1 {
		// the case executed with no args
		dir, _ := os.Getwd()
		dirs = append(dirs, dir)
	} else {
		// the case executed with args (= dirs to list)
		dirs = os.Args[1:]
	}

	for _, dir := range dirs {
		files, err := ioutil.ReadDir(dir)

		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}

		for _, file := range files {
			if isDotName(file) {
				continue
			}
			fmt.Printf("%s\t", file.Name())
		}
	}

	fmt.Printf("\n")
}
