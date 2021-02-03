package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	records "sfcard-freee-csv/records"
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

type sfViewer struct {
	day         string
	inTeiki     string
	inCorp      string
	inStation   string
	outTeiki    string
	outCorp     string
	outStation  string
	memo        string
	price       string
	remainPrice string
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

func getSfViewerList(sh *xlsx.Sheet) []sfViewer {
	var sfList []sfViewer
	isFirst := true

	sh.ForEachRow(func(r *xlsx.Row) error {
		if isFirst {
			isFirst = false
		} else if r.GetCell(9).String() != "x" {
			time, _ := r.GetCell(0).GetTime(false)
			sf := sfViewer{
				day:         time.Format("2006-01-02"),
				inTeiki:     r.GetCell(1).String(),
				inCorp:      r.GetCell(2).String(),
				inStation:   r.GetCell(3).String(),
				outTeiki:    r.GetCell(4).String(),
				outCorp:     r.GetCell(5).String(),
				outStation:  r.GetCell(6).String(),
				price:       r.GetCell(7).String(),
				remainPrice: r.GetCell(8).String(),
				memo:        r.GetCell(9).String(),
			}
			sfList = append(sfList, sf)
		}
		return nil
	}, xlsx.SkipEmptyRows)

	return sfList
}

func writeCsv(sfList []sfViewer) {

	records := getCsvData(sfList)

	f, err := os.Create("file.csv")
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	w := csv.NewWriter(f)

	w.Comma = ','
	w.UseCRLF = true

	for _, record := range records.GetCSVData() {
		if err := w.Write(record); err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
	}
	w.Flush()
}

func getCsvData(sfList []sfViewer) records.Records {

	for _, sf := range sfList {
		var r records.Record
		r.SetDay(sf.day)

		price, _ := strconv.ParseFloat(sf.price, 64)
		if price == 0 {
			continue
		} else if price > 0 {
			r.SetPrice("-" + sf.price)
		} else {
			r.SetPrice(strconv.FormatFloat(math.Abs(price), 'f', 0, 64))
		}

		r.SetRemainPrice(sf.remainPrice)
		if len(sf.inCorp) > 0 {
			body := sf.inCorp + " " + sf.inStation
			if len(sf.outCorp) > 0 {
				body = body + " - " + sf.outCorp + " " + sf.outStation
			} else {
				body = body + ":" + sf.memo
			}
			r.SetBody(body)
		} else {
			r.SetBody(sf.memo)
		}
		r.AddRecordToRecords()
	}
	return records.RecordData
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
