package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/docopt/docopt-go"
	"github.com/reconquest/pom-go"
)

var (
	version = "[manual build]"
	usage   = "pom " + version + `

Usage:
  pom [options] <key>
  pom -h | --help
  pom --version

Options:
  -i <file>  Pom file. [default: /dev/stdin]
  -h --help  Show this screen.
  --version  Show version.
`
)

func main() {
	args, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		panic(err)
	}

	contents, err := ioutil.ReadFile(args["-i"].(string))
	if err != nil {
		log.Fatal(err)
	}

	model, err := pom.Unmarshal(contents)
	if err != nil {
		log.Fatal(err)
	}

	value, err := model.Get(args["<key>"].(string))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value)
}
