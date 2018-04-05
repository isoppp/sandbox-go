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
	var (
		count   int
		flagSet = flag.NewFlagSet("default", flag.PanicOnError)
	)

	flagSet.IntVar(&count, "n", 10, "show line count")
	flagSet.Parse(os.Args[1:])

	return count, flagSet.Args()
}

func readFile(filename string, showCount int) string {
	var (
		log       = ""
		data, err = ioutil.ReadFile(filename)
	)

	if err != nil {
		fmt.Printf("error!\n%s", err)
		os.Exit(0)
	}

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

func readDir(dirPath string, filePattern string, showCount int) string {
	var (
		log            = ""
		filesInfo, err = ioutil.ReadDir(dirPath)
	)

	if err != nil {
		fmt.Println("error! ", err)
		os.Exit(0)
	}

	for _, fileInfo := range filesInfo {
		var findName = fileInfo.Name()
		if filePattern != "" {
			var match, _ = path.Match(filePattern, findName)
			if !match {
				continue
			}
		}

		var isDir, _ = isDirectory(dirPath + "/" + findName)
		if isDir {
			continue
		}
		log += readFile(dirPath+"/"+findName, showCount)
	}

	return log
}

func isDirectory(path string) (bool, error) {
	var info, err = os.Stat(path)

	if err != nil {
		fmt.Println("error! ", err)
		os.Exit(0)
	}

	return info.IsDir(), err
}

func main() {
	var (
		log              = ""
		showCount, files = getArgs()
		currentDir, _    = os.Getwd()
	)

	for _, file := range files {
		var dirName, filePattern = path.Split(file)

		if dirName == "" || dirName == "." || dirName == "./" {
			dirName = currentDir + "/"
		}

		if filePattern == "." {
			filePattern = ""
		}

		var isDir, err = isDirectory(dirName + filePattern)

		if err != nil {
			fmt.Println("error! ", err)
			os.Exit(0)
		}

		if isDir {
			dirName = dirName + filePattern
			filePattern = ""
			log += readDir(dirName, filePattern, showCount)
		} else {
			log += readFile(file, showCount)
		}
	}

	fmt.Print(log)
}
