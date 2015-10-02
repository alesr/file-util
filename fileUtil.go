package fileUtil

import (
	"io/ioutil"
	"log"
)

// ReadFile - A simple file reader.
func ReadFile(file string) string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Error on reading file.", err)
	}
	return string(data[:len(data)])
}
