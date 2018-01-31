package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"bytes"
	"strings"
	"regexp"
)

func isDotName(file os.FileInfo) bool {
	fileNameRune := []rune(file.Name())
	return fileNameRune[0] == rune('.')
}

func ls(output_buffer *bytes.Buffer, args []string) {
	var dirs []string

	args_options := make([]string, 0)

	for _, a := range args {
		option, err := regexp.MatchString("^-", a)

		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		} else if option {
			// add to the options list
			args_options = append(args_options, a)
			//fmt.Printf("%s\n", args_options)
		} else {
			// add to the files/directories list
			dirs = append(dirs, a)
		}
	}

	// parse argument
	option_all := false
	for _, o := range args_options {
		if strings.Contains(o, "a") {
			option_all = true
		}
	}

	// if no files/directories are specified, list the current directory
	if len(dirs) == 0 {
		// the case executed with no args
		dir, _ := os.Getwd()
		dirs = append(dirs, dir)
	}

	//fmt.Printf("%s\n", option_all)
	//fmt.Printf("%s\n", dirs)
	for _, dir := range dirs {
		files, err := ioutil.ReadDir(dir)

		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}

		for _, file := range files {
			//fmt.Printf("%s\n", file.Name())
			if isDotName(file) && !option_all {
				continue
			}
			output_buffer.WriteString(file.Name())
			output_buffer.WriteString("\t")
		}
	}
}

//
// main
//
func main() {
	var output_buffer bytes.Buffer

	ls(&output_buffer, os.Args[1:])

	fmt.Printf("%s\n", output_buffer.String())
}
