package filesystem

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func ListDirectory(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	fs := make([]string, 0)
	for _, f := range files {
		fs = append(fs, f.Name())
	}

	return fs
}
