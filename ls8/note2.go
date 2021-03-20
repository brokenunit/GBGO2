package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	contentHashes := make(map[string]string)
	var dir string
	fmt.Println("Укажите директорию")
	fmt.Scanln(&dir)
	if err := readDir(dir, contentHashes); err != nil {
		log.Fatal(err)
	}
}

func readDir(dirName string, contentHashes map[string]string) (err error) {
	filesInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		return
	}
	for _, fi := range filesInfos {
		if fi.IsDir() {
			err := readDir(dirName+fi.Name()+"/", contentHashes)
			if err != nil {
				return err
			}
		} else {
			// The important bits for this question
			location := dirName + fi.Name()
			// open the file
			f, err := os.Open(location)
			if err != nil {
				return err
			}
			h := md5.New()
			// copy the file body into the hash function
			if _, err := io.Copy(h, f); err != nil {
				return err
			}
			// Check if a file body with the same hash already exists
			key := fmt.Sprintf("%x", h.Sum(nil))
			if val, exists := contentHashes[key]; exists {
				fmt.Println("Duplicate found", val, location)
			} else {
				contentHashes[key] = location
			}
		}
	}
	return
}
