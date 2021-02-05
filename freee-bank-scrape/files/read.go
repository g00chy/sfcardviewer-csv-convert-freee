package files

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

type bank []string

// Banks は各データの集合
type Banks struct {
	bpro    bank
	bper    bank
	bdc     bank
	cc      bank
	service bank
}

// BanksData Banks
var BanksData Banks

func init() {

}

// Read HTML読み込み
func Read() *Banks {
	f, err := os.Open("index.html")
	if err != nil {
		panic("error")
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)

	det := chardet.NewTextDetector()
	detRslt, _ := det.DetectBest(buf)
	// fmt.Println(detRslt.Charset)

	bReader := bytes.NewReader(buf)
	reader, _ := charset.NewReaderLabel(detRslt.Charset, bReader)

	doc, _ := goquery.NewDocumentFromReader(reader)

	// title := doc.Find("title").Text()

	doc.Find("div.sync-bank-table").Each(func(i int, s *goquery.Selection) {
		s.Find("td").Each(func(_ int, s *goquery.Selection) {
			bank := s.Text()
			bank = strings.Replace(bank, " ", "", -1)
			bank = strings.Replace(bank, "\n", "", -1)
			switch i {
			case 0:
				BanksData.bpro = append(BanksData.bpro, bank)
				break
			case 1:
				BanksData.bper = append(BanksData.bper, bank)
				break
			case 2:
				BanksData.bdc = append(BanksData.bdc, bank)
				break
			case 3:
				BanksData.cc = append(BanksData.cc, bank)
				break
			case 4:
				BanksData.service = append(BanksData.service, bank)
				break
			}
		})
	})

	return &BanksData
}
