/*
	Original Author : Supan Adit Pratama
	Email : email@supanadit.com
	Current Version: 1.0
*/
package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/supanadit/gostay/helper"
	"os"
)

func main() {
	parser := argparse.NewParser("gostay", "gostay is `go get` alternative, for people who don't like `go get`")
	s := parser.String("u", "url", &argparse.Options{Required: true, Help: "URL for package"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	if *s != "" {
		data := helper.Commander{}
		data.Get(*s)
	}
}
