package test

import "io/ioutil"

// CreateBinaryFile creates a binary file with random content at the given path.
func CreateBinaryFile(path string) error {
	content := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	return ioutil.WriteFile(path, content, 0644)
}
