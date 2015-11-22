package main

import (
	"fmt"
	"strings"
)

func mergeFile(filePaths []string, startIndex, endIndex int) (filePath string) {
	if startIndex >= endIndex {
		return filePaths[startIndex]
	}
	midIndex := (startIndex + endIndex) >> 1
	frontPart := mergeFile(filePaths, startIndex, midIndex)
	backPart := mergeFile(filePaths, midIndex + 1, endIndex)
	
	return doMerge(frontPart, backPart)
}

func doMerge(frontPart, backPart string) (filePath string) {
	frontIndexs := strings.Split(strings.Split(frontPart, "-")[1], ".")
	startIndex := frontIndexs[0]
	
	backIndexs := strings.Split(strings.Split(backPart, "-")[1], ".")
	endIndex := backIndexs[len(backIndexs) - 1]
	
	newFileName = "tmpFile-" + startIndex + "." + endIndex
	
	fmt.Println("merge file (" + startIndex + "-" + endIndex + ")")
	return newFileName;
}

func splitFile(input string) (filePaths []string) {
	result := []string{"tmpfile-1", "tmpfile-2", "tmpfile-3", "tmpfile-4"}
	return result
}

func main() {
	filePaths := splitFile("input")
	finalPath := mergeFile(filePaths, 0, 3)
	fmt.Println(finalPath)
}