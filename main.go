package main

import (
	"flag"
	"fmt"
	"os"
)

const MaxLen = 50

type Args struct {
	useGit    bool
	allowStr  string
	ignoreStr string
	ignoreDir string
}

var args = Args{}

func parseArgs() {
	flag.BoolVar(&args.useGit, "useGit", false, "是否使用gitignore规则")
	flag.StringVar(&args.allowStr, "allowExt", "*", `允许的文件后缀名，存在多个时请使用""包裹如："py go js" 默认全匹配`)
	flag.StringVar(&args.ignoreStr, "ignoreExt", "", `忽略的文件后缀名，存在多个时请使用""包裹如："py go js" 默认不忽略`)
	flag.StringVar(&args.ignoreDir, "ignoreDir", "", `允许的文件夹名，存在多个时请使用""包裹如："node_modules" 默认不忽略`)
	flag.Parse()
}

func main() {
	parseArgs()
	rootDir := os.Args[len(os.Args)-1]

	var rule Rule
	if args.useGit {
		rule = &gitRule{}
	} else {
		rule = &samRule{
			AllowExt:  parseStringToMylist(args.allowStr),
			IgnoreExt: parseStringToMylist(args.ignoreStr),
			IgnoreDir: parseStringToMylist(args.ignoreDir),
		}
	}
	rule.init()
	count := Count{
		rule:      rule,
		mainCount: make(map[string]int),
		fileCount: make(map[string]int),
	}
	if len(os.Args) < 2 {
		fmt.Println("请输入文件夹")
		return
	}
	if !IsDir(rootDir) {
		fmt.Println("找不到该文件夹！")
		return
	}
	all := count.countDir(rootDir)
	render(count.mainCount, "文件类型", all)
	render(count.fileCount, "文件名", all)
}
