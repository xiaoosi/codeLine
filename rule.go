package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type myList []string

func (ml myList) has(s string) bool {
	for _, item := range ml {
		if item == s {
			return true
		}
	}
	return false
}

type Rule interface {
	init(rootPath string)
	isFilePass(path string, file os.FileInfo) bool
	isDirPass(pathName string) bool
}

// sample Rule
type samRule struct {
	AllowExt  myList
	IgnoreExt myList
	IgnoreDir myList
	rootPath  string
}

func (r *samRule) init(rootPath string) {
	r.rootPath = rootPath
}

func getExt(file os.FileInfo) string {
	name := file.Name()
	strList := strings.Split(name, ".")
	if len(strList) > 1 {
		return strList[1]
	}
	return ""
}

func (r *samRule) isFilePass(path string, file os.FileInfo) bool {
	ext := getExt(file)
	if !r.AllowExt.has("*") && !r.AllowExt.has(ext) {
		return false
	}
	if r.IgnoreExt.has(ext) {
		return false
	}
	return true
}

func (r *samRule) isDirPass(pathName string) bool {
	if r.IgnoreDir.has(pathName) {
		return false
	}
	return true
}

func parseStringToMylist(str string) myList {
	if str == "" {
		return myList{}
	}
	res := myList{}
	items := strings.Split(str, " ")
	for _, item := range items {
		if item != "" {
			res = append(res, item)
		}
	}
	return res
}

// gitignore Role
type gitRule struct {
	paths    myList
	rootPath string
}

func (r *gitRule) init(rootPath string) {
	if !strings.HasSuffix(rootPath, "/"){
		rootPath += "/"
	}
	r.rootPath = rootPath
	cmd := exec.Command("git", "version")
	strBytes, err := cmd.Output()
	if err != nil {
		log.Fatalf("err: %s\ncon't run git in your os", err.Error())
	}
	str := string(strBytes)
	if strings.Trim(str, "\n ") == "" {
		log.Fatalf("err: %s\ncon't find git in your os")
	}
	cmd = exec.Command("git", "--git-dir="+rootPath+".git", "ls-files")
	strBytes, err = cmd.Output()
	if err != nil {
		log.Fatalf("err: %s\nPlease make sure git is initialized\ntry run `git init && git add .`", err.Error())
	}
	str = string(strBytes)
	fileList := strings.Split(str, "\n")
	if len(fileList) > 0 {
		fileList = fileList[:len(fileList)-1]
	}
	newFileList := make([]string, 0)
	for _, fileStr := range fileList {
		newFileList = append(newFileList, filepath.Join(rootPath, fileStr))
	}
	r.paths = newFileList
}

func (r *gitRule) isFilePass(path string, file os.FileInfo) bool {
	fullPath := filepath.Join(path, file.Name())
	if r.paths.has(fullPath) {
		return true
	}
	return false
}

func (r *gitRule) isDirPass(pathName string) bool {
	return true
}
