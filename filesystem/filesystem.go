package filesystem

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func GetCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func ListDirectory(path string) ([]string, error) {
	if !DoesDirectoryExist(path) {
		return nil, errors.New(fmt.Sprintf("no such directory: %s", path))
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	fs := make([]string, 0)
	for _, f := range files {
		if f.IsDir() {
			fs = append(fs, fmt.Sprintf("%s/", f.Name()))
		} else {
			fs = append(fs, f.Name())
		}
	}

	return fs, nil
}

func DoesDirectoryExist(path string) bool {
	f, err := os.Stat(path)
	return !os.IsNotExist(err) && f.IsDir()
}

func DoesFileExist(path string) bool {
	f, err := os.Stat(path)
	return !os.IsNotExist(err) && !f.IsDir()
}

func MakeAbsolute(path string) string {
	p, err := filepath.Abs(path)

	if err != nil {
		log.Fatal(err)
	}

	return p
}

func IsAbsolute(path string) bool {
	return string(path[0]) == "/"
}
