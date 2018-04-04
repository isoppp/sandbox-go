package main

import (
	"fmt"
	"os"
	"flag"
	"io/ioutil"
	"strings"
)

var flagSet = flag.NewFlagSet("default", flag.PanicOnError)

func getArgs() (int, []string) {
	var count int
	flagSet.IntVar(&count, "n", 10, "help message")
	flagSet.Parse(os.Args[1:])
	fmt.Println(flagSet.Args())

	return count, flagSet.Args()
}

func readFile(filename string, showCount int) string {
	log := ""
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Printf("error!\n%s", err)
		os.Exit(0)
	}

	log += "=== " + filename + " ===\n"

	for i, line := range strings.Split(string(data), "\n") {
		if i == showCount {
			break
		}
		log += line + "\n"
	}

	return log
}

func main() {
	log := ""
	showCount, files := getArgs()

	for _, file := range files {
		log += readFile(file, showCount)
	}

	fmt.Print(log)
}
