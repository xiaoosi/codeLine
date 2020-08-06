# CodeLine 
一个代码行统计小工具
* 支持自定义规则
* .gitignore

### 下载

### 使用
基础用法
```
codeLine --useGit .
┌───────────┬──────┬──────┐
│ 文件类型  │ 行数 │ 占比 │
├───────────┼──────┼──────┤
│ go        │  318 │ 0.81 │
│ sum       │   33 │ 0.08 │
│ MD        │   29 │ 0.07 │
│ mod       │   10 │ 0.03 │
│ gitignore │    1 │ 0.00 │
├───────────┼──────┼──────┤
│ 总行数    │  391 │      │
└───────────┴──────┴──────┘
┌────────────┬──────┬──────┐
│ 文件名     │ 行数 │ 占比 │
├────────────┼──────┼──────┤
│ rule.go    │  118 │ 0.30 │
│ count.go   │   74 │ 0.19 │
│ main.go    │   59 │ 0.15 │
│ sortMap.go │   34 │ 0.09 │
│ render.go  │   33 │ 0.08 │
│ go.sum     │   33 │ 0.08 │
│ ReadMe.MD  │   29 │ 0.07 │
│ go.mod     │   10 │ 0.03 │
│ .gitignore │    1 │ 0.00 │
├────────────┼──────┼──────┤
│ 总行数     │  391 │      │
└────────────┴──────┴──────┘

```
### 参数
```
Usage of codeLine:
  -allowExt string
        允许的文件后缀名，存在多个时请使用""包裹。如："py go js" 默认全匹配 (default "*")
  -ignoreDir string
        允许的文件夹名，存在多个时请使用""包裹。如："node_modules" 默认不忽略
  -ignoreExt string
        忽略的文件后缀名，存在多个时请使用""包裹。如："py go js" 默认不忽略
  -useGit
        是否使用gitignore规则

```

### 依赖
[go-pretty](https://github.com/jedib0t/go-pretty/tree/master/table)