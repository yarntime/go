package utils

import (
	"encoding/csv"
	"github.com/golang/glog"
	"io"
	"os"
)

type CsvUtil struct {
	fileName string
	reader   *csv.Reader
	file     *os.File
}

func NewCsvUtil(name string) *CsvUtil {
	return &CsvUtil{
		fileName: name,
	}
}

func (c *CsvUtil) Open() {
	file, err := os.Open(c.fileName)
	if err != nil {
		glog.Fatalf("Failed to open file %s: %s", c.fileName, err.Error())
	}
	c.file = file
	c.reader = csv.NewReader(file)
}

func (c *CsvUtil) Close() {
	c.file.Close()
}

func (c *CsvUtil) Read() ([]string, error) {
	record, err := c.reader.Read()
	if err == io.EOF {
		return []string{}, err
	} else if err != nil {
		return []string{}, err
	}
	return record, nil
}
