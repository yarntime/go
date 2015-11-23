package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"strconv"
	"io"
)

func mergeFile(filePaths [1001]string, startIndex, endIndex int) (filePath string) {
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
	
	newFileName := "tmpFile-" + startIndex + "." + endIndex
	
	fmt.Println("merge file (" + startIndex + "-" + endIndex + ")")
	return newFileName;
}

func splitFile(input string) (filePaths [1001]string) {
	//result := []string{"tmpfile-1", "tmpfile-2", "tmpfile-3", "tmpfile-4"}
	result := [1001]string{}
	inputFile, inputError := os.OpenFile(input, os.O_RDONLY, 0660)
	if inputError != nil { panic(inputError) }
	defer inputFile.Close()
	fileIndex, curLine := 1, 0
	filePath := "tmpfile-"
	buf := [10001]int{}
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			break
		}
		buf[curLine], _ = strconv.Atoi(inputString)
		curLine++
		if curLine == 10000 {
			// sort file here
			tmpFilePath := filePath + strconv.Itoa(fileIndex)
			flushToFile(tmpFilePath, buf, curLine, os.O_CREATE|os.O_RDWR)
			result[fileIndex] = tmpFilePath
			fileIndex++
			curLine = 0
		}
	}
	tmpFilePath := filePath + strconv.Itoa(fileIndex)
	flushToFile(tmpFilePath, buf, curLine, os.O_CREATE|os.O_RDWR)
	result[fileIndex] = tmpFilePath
	return result
}

func flushToFile(filePath string, buf [10001]int, length, model int) {
	inputFile, inputError := os.OpenFile(filePath, model, 0660)
	if inputError != nil {
		panic(inputError)
	}
	defer inputFile.Close()
	dataWrite := bufio.NewWriter(inputFile)
	
	for i := 0; i < length; i++ {
		dataWrite.WriteString(strconv.Itoa(buf[i]) + "\n")
	}
}

func main() {
	filePaths := splitFile("input.txt")
	finalPath := mergeFile(filePaths, 0, 3)
	fmt.Println(finalPath)
}