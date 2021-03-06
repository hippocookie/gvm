package classpath

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	fmt.Printf("file:%s\n", fileName)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}
