package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Count struct {
	rule      Rule
	mainCount map[string]int
	fileCount map[string]int
}

func (c *Count) init() {
	c.mainCount = make(map[string]int)
	c.fileCount = make(map[string]int)

}

func (c *Count) countFile(path string, file os.FileInfo) int {
	res := 0
	if !file.IsDir() && c.rule.isFilePass(path, file) {
		fullPath := filepath.Join(path, file.Name())
		fi, err := os.Open(fullPath)
		if err != nil {
			log.Println("read error:", err)
			return 0
		}
		defer fi.Close()
		br := bufio.NewReader(fi)
		for {
			_, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			res++
		}
		c.mainCount[getExt(file)] += res
		c.fileCount[fullPath] += res
	}
	return res
}
func (c *Count) countDir(path string) int {
	res := 0
	if !c.rule.isDirPass(path) {
		return 0
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println("con't find this dir")
		return 0
	}
	for _, file := range files {
		if file.IsDir() {
			res += c.countDir(filepath.Join(path, file.Name()))
		} else {
			res += c.countFile(path, file)
		}
	}
	return res
}


func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
