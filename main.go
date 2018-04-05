package main

import (
	"fmt"
	"os"
	"flag"
	"io/ioutil"
	"strings"
	"path"
)

func getArgs() (int, []string) {
	var count int
	var flagSet = flag.NewFlagSet("default", flag.PanicOnError)
	flagSet.IntVar(&count, "n", 10, "show line count")
	flagSet.Parse(os.Args[1:])

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

func readDir(dirPath string, filePattern string, showCount int) string {
	log := ""
	fileInfos, err := ioutil.ReadDir(dirPath)

	if err != nil {
		fmt.Println("error! ", err)
		os.Exit(0)
	}

	for _, fileInfo := range fileInfos {
		var findName = fileInfo.Name()
		var matched = true
		if filePattern != "" {
			matched, _ = path.Match(filePattern, findName)
		}

		if !matched {
			continue
		}

		log += readFile(dirPath+findName, showCount)
	}

	return log
}

func isDirectory(path string) (bool, error) {
	var info, err = os.Stat(path)

	if err != nil {
		return false, err
	}

	return info.IsDir(), err
}

func main() {
	log := ""
	var showCount, files = getArgs()
	var currentDir, _ = os.Getwd()

	for _, file := range files {
		var dirName, filePattern = path.Split(file)
		if dirName == "" {
			dirName = currentDir + "/"
		}
		fmt.Println(dirName, filePattern)

		var isDir, err = isDirectory(dirName + filePattern)
		if err != nil {
			fmt.Printf("error!\n%s", err)
			os.Exit(0)
		}
		if isDir {
			dirName = dirName + filePattern
			filePattern = ""
		}

		fmt.Println("last " + dirName)

		if isDir {
			log += readDir(dirName, filePattern, showCount)
		} else {
			log += readFile(file, showCount)
		}
	}

	fmt.Print(log)
}
