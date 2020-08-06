# CodeLine 
一个代码行统计小工具
* 支持自定义规则
* 支持.gitignore过滤规则

### 下载
[release1.0](https://github.com/xiaoosi/codeLine/releases/tag/v1.0)
### 使用
基础用法
```bash
codeLine --useGit .
```
```
┌───────────┬───────┬────────────┐
│ FILE TYPE │ LINES │ PROPORTION │
├───────────┼───────┼────────────┤
│ go        │   318 │ 0.76       │
│ MD        │    55 │ 0.13       │
│ sum       │    33 │ 0.08       │
│ mod       │    10 │ 0.02       │
│ gitignore │     1 │ 0.00       │
├───────────┼───────┼────────────┤
│ ALL       │   417 │            │
└───────────┴───────┴────────────┘
┌────────────┬───────┬────────────┐
│ FILE PATH  │ LINES │ PROPORTION │
├────────────┼───────┼────────────┤
│ rule.go    │   118 │ 0.28       │
│ count.go   │    74 │ 0.18       │
│ main.go    │    59 │ 0.14       │
│ ReadMe.MD  │    55 │ 0.13       │
│ sortMap.go │    34 │ 0.08       │
│ go.sum     │    33 │ 0.08       │
│ render.go  │    33 │ 0.08       │
│ go.mod     │    10 │ 0.02       │
│ .gitignore │     1 │ 0.00       │
├────────────┼───────┼────────────┤
│ ALL        │   417 │            │
└────────────┴───────┴────────────┘

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