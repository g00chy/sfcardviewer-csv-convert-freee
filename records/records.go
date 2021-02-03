package records

import "sort"

type Header struct {
	head []string
}
type Record struct {
	day         string
	price       string
	remainPrice string
	body        string
}

type Records []Record

var head []string
var RecordData Records

func init() {
	head = []string{"取引日", "金額", "残高", "取引内容"}
}

func (r *Record) SetDay(day string) {
	r.day = day
}

func (r *Record) SetPrice(price string) {
	r.price = price
}

func (rs *Records) GetCSVData() [][]string {
	reverse()

	var retData [][]string

	retData = append(retData, head)
	for _, record := range RecordData {
		retData = append(retData, []string{record.day, record.price, record.remainPrice, record.body})
	}
	return retData
}

func (r *Record) SetRemainPrice(remainPrice string) {
	r.remainPrice = remainPrice
}

func (r *Record) SetBody(body string) {
	r.body = body
}

func (r *Record) AddRecordToRecords() {
	RecordData = append(RecordData, *r)
}

func (p Records) Len() int {
	return len(p)
}

func (p Records) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Records) Less(i, j int) bool {
	return i < j
}

func reverse() {
	sort.SliceStable(RecordData, func(i, j int) bool {
		return i > j
	})
}
