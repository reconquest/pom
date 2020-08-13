package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/docopt/docopt-go"
	"github.com/reconquest/pom"
)

var (
	version = "[manual build]"
	usage   = "pom " + version + `

Usage:
  pom [options] [-p] <key> [<value>]
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

	propertyMode := args["-p"].(bool)
	key := args["<key>"].(string)
	value, withValue := args["<value>"].(string)

	contents, err := ioutil.ReadFile(args["-i"].(string))
	if err != nil {
		log.Fatal(err)
	}

	model, err := pom.Unmarshal(contents)
	if err != nil {
		log.Fatal(err)
	}

	if !withValue {
		var actual string
		if propertyMode {
			actual, err = model.GetProperty(key)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			actual, err = model.Get(key)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println(actual)

		return
	}

	if propertyMode {
		model.SetProperty(key, value)
		data, err := model.Marshal()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(data))
	} else {
		panic("Not implemented")
	}
}
