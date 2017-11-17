package TDUBus

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func GetTimeSchedule() (result Schedule) {
	url := "https://www.dendai.ac.jp/access/saitama_hatoyama/school_bus.html"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Print("url scarapping failed")
	}
	//	s := doc.Find("h4")
	//	fmt.Println()
	// DOMのまま出力
	doc.Find("h4").Each(func(n int, s *goquery.Selection) {
		tablename := strings.TrimSpace(s.Text())
		if tablename == "大学発" {
			data := s.Next()
			header := GetHeader(data)
			fmt.Println(header)
			result = GetTime(data)
		}
	})
	return
}

type Schedule struct {
	Takasaka   map[string][]string
	Kitasakado map[string][]string
	Kumagaya   map[string][]string
	Kounosu    map[string][]string
}

func GetTime(data *goquery.Selection) Schedule {
	ret := Schedule{}
	ret.Takasaka = make(map[string][]string)
	ret.Kitasakado = make(map[string][]string)
	ret.Kumagaya = make(map[string][]string)
	ret.Kounosu = make(map[string][]string)

	d := data.Find("tr").Slice(1, 16)
	d.Each(func(n int, s *goquery.Selection) {
		hour := strings.TrimSpace(s.Find("th").Text())
		mbase := s.Find("td")
		ret.Takasaka[hour] = MakeSchedule(mbase.Slice(0, 1))
		ret.Kitasakado[hour] = MakeSchedule(mbase.Slice(1, 2))
		ret.Kumagaya[hour] = MakeSchedule(mbase.Slice(2, 3))
		ret.Kounosu[hour] = MakeSchedule(mbase.Slice(3, 4))
	})
	return ret
}

func MakeSchedule(data *goquery.Selection) []string {
	ret := []string{}
	d := strings.TrimSpace(data.Text())
	ret = strings.Fields(d)
	return ret
}

func GetHeader(data *goquery.Selection) []string {
	ret := []string{}
	d := data.Find("tr > th").Slice(1, 5)
	d.Each(func(n int, s *goquery.Selection) {
		tablename := strings.TrimSpace(s.Text())
		ret = append(ret, tablename)
	})
	return ret
}
