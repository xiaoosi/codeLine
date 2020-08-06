package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//rule := samRule{
	//	AllowExt:  myList{"*"},
	//	IgnoreExt: myList{},
	//	IgnoreDir: myList{},
	//}
	rule := &gitRule{}
	rule.init()
	count := Count{
		rule:      rule,
		mainCount: make(map[string]int),
		fileCount: make(map[string]int),
	}
	if len(os.Args) < 2 {
		log.Fatalln("输入文件夹参数！")
	}
	path := os.Args[1]
	if !IsDir(path) {
		log.Fatalln("找不到该文件夹！")
	}
	fmt.Println(count.countDir(path))
	fmt.Println(count.mainCount)
	fmt.Println(count.fileCount)

}
