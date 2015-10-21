package fileUtil

import (
	"io/ioutil"
	"os/user"
)

// ReadFile - A simple file reader.
func ReadFile(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	return string(data[:len(data)]), err
}

// FindUserHomeDir outputs the filepath for home directory.
func FindUserHomeDir() (string, error) {
	usr, err := user.Current()
	return usr.HomeDir, err
}
