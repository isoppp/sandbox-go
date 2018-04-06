package main

import (
	"fmt"
	"os"
	"flag"
	"io/ioutil"
	"strings"
	"path/filepath"
)

func judgeError(error error) {
	if error != nil {
		fmt.Printf("error!\n%s", error)
		os.Exit(0)
	}
}

type Option struct {
	searchSubtree bool
	count         int
}

func getArgs() (Option, []string) {
	var (
		count   int
		subtree bool
		flagSet = flag.NewFlagSet("default", flag.PanicOnError)
	)

	flagSet.IntVar(&count, "n", 10, "Number of rows to display")
	flagSet.BoolVar(&subtree, "s", false, "If you want to search directories")
	flagSet.Parse(os.Args[1:])

	return Option{count: count, searchSubtree: subtree}, flagSet.Args()
}

func readFile(filename string, showCount int) string {
	var (
		log       = ""
		data, err = ioutil.ReadFile(filename)
	)

	judgeError(err)

	log += "=== " + filename + " ===\n"

	for i, line := range strings.Split(string(data), "\n") {
		if i == showCount {
			log += "\n"
			break
		}
		log += line + "\n"
	}

	return log
}

func readDir(dirPath string, option Option) string {
	var (
		log            = ""
		filesInfo, err = ioutil.ReadDir(dirPath)
	)

	judgeError(err)

	for _, info := range filesInfo {
		var findPath = filepath.Join(dirPath + "/" + info.Name())
		var isDir = isDirectory(findPath)

		if isDir {
			if option.searchSubtree {
				log += readDir(findPath, option)
			} else {
				fmt.Println(findPath + " is directory. If you want search directories, use -s option")
			}
		} else {
			log += readFile(findPath, option.count)
		}
	}

	return log
}

func isDirectory(path string) bool {
	var info, err = os.Stat(path)

	judgeError(err)

	return info.IsDir()
}

func main() {
	var (
		log           = ""
		option, files = getArgs()
	)

	for _, file := range files {
		var absPath, err = filepath.Abs(file)
		judgeError(err)

		if isDirectory(absPath) {
			if option.searchSubtree {
				log += readDir(absPath, option)
			} else {
				fmt.Println(absPath + " is directory. If you want search directories, use -s option")
				continue;
			}
		} else {
			log += readFile(absPath, option.count)
		}
	}

	fmt.Print(log)
}
