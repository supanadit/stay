/*
	Original Author : Supan Adit Pratama
	Email : email@supanadit.com
	Current Version: 1.3
*/
package main

import (
	"bufio"
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/supanadit/stay/helper"
	"log"
	"os"
)

func main() {
	parser := argparse.NewParser("stay", "Bulletproof package manager")
	file := parser.String("f", "file", &argparse.Options{Required: false, Help: "File target like `requirement.txt`"})
	u := parser.String("u", "url", &argparse.Options{Required: false, Help: "URL for package"})
	a := parser.String("a", "get-related", &argparse.Options{Required: false, Help: "Get Package and find related package"})
	r := parser.String("r", "remove", &argparse.Options{Required: false, Help: "Remove installed package"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	if *u != "" {
		data := helper.Commander{}
		data.Get(*u, false)
	}

	if *a != "" {
		data := helper.Commander{}
		data.Get(*a, true)
	}

	if *r != "" {
		data := helper.Commander{}
		data.Delete(*r)
	}

	// For Automatically Process
	if *file != "" {
		file, err := os.Open(*file)

		if err != nil {
			log.Fatalf("Failed opening file: %u", err)
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var txtlines []string

		for scanner.Scan() {
			txtlines = append(txtlines, scanner.Text())
		}

		_ = file.Close()

		for _, eachline := range txtlines {
			data := helper.Commander{}
			fmt.Printf("%u ", eachline)
			data.Get(eachline, false)
		}
	}

	if *u == "" && *file == "" && *a == "" && *r == "" {
		fmt.Print(parser.Usage(err))
	}
}
