package main

import (
	"fmt"
	"os"
	"flag"
	"io/ioutil"
	"strings"
)

var flagSet = flag.NewFlagSet("default", flag.PanicOnError)

func getArgs() (int, string) {
	var count int
	flagSet.IntVar(&count, "n", 10, "help message")
	flagSet.Parse(os.Args[1:])

	return count, flagSet.Args()[0]
}

func main() {
	log := ""
	showCount, fileName := getArgs()
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Printf("error!\n%s", err)
		os.Exit(0)
	}

	log += "=== " + fileName + " ===\n"

	for i, line := range strings.Split(string(data), "\n") {
		if i == showCount {
			break
		}
		log += line + "\n"
	}
	fmt.Print(log)
}
