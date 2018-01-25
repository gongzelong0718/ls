package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"bytes"
)

func isDotName(file os.FileInfo) bool {
	fileNameRune := []rune(file.Name())
	return fileNameRune[0] == rune('.')
}

func ls(output_buffer *bytes.Buffer, args []string) {
	var dirs []string

	if len(os.Args) == 1 {
		// the case executed with no args
		dir, _ := os.Getwd()
		dirs = append(dirs, dir)
	} else {
		// the case executed with args (= dirs to list)
		dirs = args
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
			//fmt.Printf("%s\t", file.Name())
			output_buffer.WriteString(file.Name())
			output_buffer.WriteString("\t")
		}
	}

	//fmt.Printf("\n")
}

//
// main
//
func main() {
	var output_buffer bytes.Buffer

	ls(&output_buffer, os.Args[1:])

	fmt.Printf("%s\n", output_buffer.String())
}
