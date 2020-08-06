package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"os"
	"sort"
)

func convMap(strMap map[string]int, all int) []table.Row {
	items := convTo(strMap)
	sort.Sort(sort.Reverse(items))
	res := make([]table.Row, 0)
	for _, item := range items {
		// 截取过长的文件名
		k := item.key
		if len(k) > MaxLen {
			k = "..." + k[len(k)-MaxLen:]
		}
		res = append(res, table.Row{k, item.value, fmt.Sprintf("%.2f", float64(item.value)/float64(all))})
	}
	return res
}

func render(in map[string]int, dataType string, all int) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{dataType, "lines", "proportion"})
	t.AppendRows(convMap(in, all))
	t.AppendFooter(table.Row{"all", all})
	t.SetStyle(table.StyleLight)
	t.Render()
}

