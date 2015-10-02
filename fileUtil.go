package fileUtil

import (
	"io/ioutil"
	"log"
	"os/user"
)

// ReadFile - A simple file reader.
func ReadFile(file string) string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Error on reading file.", err)
	}
	return string(data[:len(data)])
}

// FindUserHomeDir outputs the filepath for home directory.
func FindUserHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal("Error on searching for home directory.", err)
	}
	return usr.HomeDir
}
