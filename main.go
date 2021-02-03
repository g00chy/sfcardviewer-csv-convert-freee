package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/tealeg/xlsx/v3"
)

type transaction struct {
	day  time.Time
	out  int
	in   int
	body string
}

type SfViewer struct {
	day        string
	inTeiki    string
	inCorp     string
	inStation  string
	outTeiki   string
	outCorp    string
	outStation string
	memo       string
	price      string
}

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

func getSfViewerList(sh *xlsx.Sheet) []SfViewer {
	var sfList []SfViewer

	sh.ForEachRow(func(r *xlsx.Row) error {
		sf := SfViewer{
			day:        r.GetCell(0).String(),
			inTeiki:    r.GetCell(1).String(),
			inCorp:     r.GetCell(2).String(),
			inStation:  r.GetCell(3).String(),
			outTeiki:   r.GetCell(4).String(),
			outCorp:    r.GetCell(5).String(),
			outStation: r.GetCell(6).String(),
			price:      r.GetCell(7).String(),
			memo:       r.GetCell(8).String(),
		}
		sfList = append(sfList, sf)
		return nil
	})

	return sfList
}

func writeCsv(sfList []SfViewer) {

	records := getCsvData(sfList)

	f, err := os.Create("file.csv")
	if err != nil {
		log.Fatal(err)
	}

	w := csv.NewWriter(f)

	w.Comma = ','    // デフォルトはカンマ区切りで出力される。変更する場合はこの rune 文字を変更する
	w.UseCRLF = true // 改行文字を CRLF(\r\n) にする

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatal(err)
		}
	}
}

func getCsvData(sfList []SfViewer) [][]string {
	records := [][]string{
		{"取引日", "出金額", "入金額", "取引内容"},
	}

	for _, sf := range sfList {
		var record []string
		record = append(record, sf.day)

		price, _ := strconv.ParseFloat(sf.price, 64)
		if price == 0 {
			continue
		} else if price > 0 {
			record = append(record, sf.price)
			record = append(record, "")
		} else {
			record = append(record, "")
			record = append(record, strconv.FormatFloat(math.Abs(price), 'f', 0, 64))
		}

		if len(sf.inCorp) > 0 {
			body := sf.inCorp + " " + sf.inStation
			if len(sf.outCorp) > 0 {
				body = body + " - " + sf.outCorp + " " + sf.outStation
			}
			record = append(record, body)
		} else {
			record = append(record, sf.memo)
		}
		records = append(records, record)
	}

	return records
}

func main() {
	sh, ok := readSheet("SUICA")

	if ok {
		sfList := getSfViewerList(sh)
		writeCsv(sfList)

	} else {
		panic("not readable")
	}
}
