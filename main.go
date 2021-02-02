package main

import (
	"fmt"

	"github.com/tealeg/xlsx/v3"
)

func readSheet(sheetName string) (*xlsx.Sheet, bool) {
	wb, err := xlsx.OpenFile("./SUICA.xlsx")
	if err != nil {
		panic(err)
	}
	sh, ok := wb.Sheet[sheetName]
	if !ok {
		fmt.Println("Sheet does not exist")
		return sh, ok
	}
	fmt.Println("Max row in sheet:", sh.MaxRow)

	return sh, ok
}

func formatCSV(sh *xlsx.Sheet) {
	row, err := sh.Row(1)
	if err != nil {
		panic(err)
	}

	sh.ForEachRow(readRow)
}

func readRow(r *xlsx.Row) err {
	c := r.GetCell(0)
	return 
}

func main() {
	sh, ok := readSheet("SUICA")

	if ok {
		formatCSV(sh)
	} else {
		panic("not readable")
	}
}

//
//var doc ods.Doc
//if err := f.ParseContent(&doc); err != nil {
//    panic(err)
//}
//
//for _, sheet := range doc.Table {
//    fmt.Print(sheet.Name)
//}
