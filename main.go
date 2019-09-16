/*
	Original Author : Supan Adit Pratama
	Email : email@supanadit.com
	Current Version: 1.0
*/
package main

import (
	"bufio"
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/supanadit/gostay/helper"
	"log"
	"os"
)

func main() {
	parser := argparse.NewParser("gostay", "Gostay is `go get` alternative, and a package manager for Golang")
	file := parser.String("f", "file", &argparse.Options{Required: false, Help: "File target like `requirement.txt`", Default: ""})
	s := parser.String("u", "url", &argparse.Options{Required: false, Help: "URL for package", Default: ""})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	if *s != "" {
		data := helper.Commander{}
		data.Get(*s)
	}

	// For Automatically Process
	if *file != "" {
		file, err := os.Open(*file)

		if err != nil {
			log.Fatalf("Failed opening file: %s", err)
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var txtlines []string

		for scanner.Scan() {
			txtlines = append(txtlines, scanner.Text())
		}

		file.Close()

		for _, eachline := range txtlines {
			data := helper.Commander{}
			fmt.Printf("%s ", eachline)
			data.Get(eachline)
		}
	}

	if *s == "" && *file == "" {
		fmt.Print(parser.Usage(err))
	}
}
