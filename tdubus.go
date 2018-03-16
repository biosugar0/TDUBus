package TDUBus

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/yut-kt/goholiday"
	"strconv"
	"strings"
	"time"
)

type BusTimes struct {
	Up   Schedule
	Down Schedule
}
type Schedule struct {
	Takasaka   map[string][]string
	Kitasakado map[string][]string
	Kumagaya   map[string][]string
	Kounosu    map[string][]string
}

type Cli struct {
}

func (c *Cli) NextDown(station string) ([]string, []string) {
	//	timetable := GetTimeSchedules()
	// 祝日は判定せずに、両方出す。
	timetables := GetTimeSchedules()
	now := time.Now()
	day := fmt.Sprint(now.Weekday())
	result := []string{}
	vacationtime := timetables[2].Down
	vacationresult := next(vacationtime, now, station)
	if goholiday.IsHoliday(now) {
		return nil, nil
	} else if day == "Saturday" {
		timetable := timetables[1].Down
		result = next(timetable, now, station)
	} else {
		timetable := timetables[0].Down
		result = next(timetable, now, station)
	}
	return result, vacationresult
}

func getnextTime(result []string, times map[string][]string, now time.Time) []string {
	for {
		hour := fmt.Sprint(now.Hour())
		min := int(now.Minute())
		for _, value := range times[hour] {
			b, _ := strconv.Atoi(value)
			if b > min {
				result = append(result, fmt.Sprintf("%v時%v分", hour, b))
			}
			if len(result) == 3 {
				return result
			}
		}
		if hour == "21" {
			return result
		}
		if len(result) < 3 {
			now = now.Add(time.Duration(1) * time.Hour)
		}
	}
}

func next(timetable Schedule, now time.Time, station string) []string {
	result := []string{}
	if station == "takasaka" {
		bustime := timetable.Takasaka
		result = getnextTime(result, bustime, now)
	} else if station == "kitasakado" {
		bustime := timetable.Kitasakado
		result = getnextTime(result, bustime, now)
	} else if station == "kumagaya" {
		bustime := timetable.Kumagaya
		result = getnextTime(result, bustime, now)
	} else if station == "kounosu" {
		bustime := timetable.Kounosu
		result = getnextTime(result, bustime, now)
	}
	return result
}

func GetTimeSchedules() []BusTimes {
	url := "https://www.dendai.ac.jp/access/saitama_hatoyama/school_bus.html"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Print("url scarapping failed")
	}

	schedules := getschedule(doc)
	result := []BusTimes{}
	tmp := BusTimes{}
	for i, d := range schedules {
		if i%2 == 0 {
			tmp = BusTimes{}
			tmp.Up = d
		} else {
			tmp.Down = d
			result = append(result, tmp)
		}
	}
	return result
}

func getschedule(doc *goquery.Document) []Schedule {
	result := []Schedule{}
	doc.Find("tbody").Each(func(i int, s *goquery.Selection) {
		if i < 2 {
			cal := getTime(s)
			result = append(result, cal)
		} else {
			cal1, cal2 := getTimeholiday(s)
			result = append(result, cal1)
			result = append(result, cal2)
		}
	})
	return result
}

func getTime(data *goquery.Selection) Schedule {
	ret := Schedule{}
	ret.Takasaka = make(map[string][]string)
	ret.Kitasakado = make(map[string][]string)
	ret.Kumagaya = make(map[string][]string)
	ret.Kounosu = make(map[string][]string)

	d := data.Find("tr").Slice(1, 16)
	d.Each(func(n int, s *goquery.Selection) {
		hour := strings.TrimSpace(s.Find("th").Text())
		mbase := s.Find("td")
		ret.Takasaka[hour] = timesplit(mbase.Slice(0, 1))
		ret.Kitasakado[hour] = timesplit(mbase.Slice(1, 2))
		ret.Kumagaya[hour] = timesplit(mbase.Slice(2, 3))
		ret.Kounosu[hour] = timesplit(mbase.Slice(3, 4))
	})
	return ret
}

func getTimeholiday(data *goquery.Selection) (Schedule, Schedule) {
	ret1 := Schedule{}
	ret2 := Schedule{}
	ret1.Takasaka = make(map[string][]string)
	ret1.Kitasakado = make(map[string][]string)
	ret1.Kumagaya = make(map[string][]string)
	ret2.Takasaka = make(map[string][]string)
	ret2.Kitasakado = make(map[string][]string)
	ret2.Kumagaya = make(map[string][]string)

	d := data.Find("tr").Slice(1, 16)
	d.Each(func(n int, s *goquery.Selection) {
		if n > 0 {
			hour := s.Find("th").Text()
			mbase := s.Find("td")
			ret1.Kumagaya[hour] = timesplit(getvalue(mbase, 0))
			ret1.Kitasakado[hour] = timesplit(getvalue(mbase, 1))
			ret1.Takasaka[hour] = timesplit(getvalue(mbase, 2))
			ret2.Takasaka[hour] = timesplit(getvalue(mbase, 3))
			ret2.Kitasakado[hour] = timesplit(getvalue(mbase, 4))
			ret2.Kumagaya[hour] = timesplit(getvalue(mbase, 5))
		}
	})
	return ret1, ret2
}

func getvalue(data *goquery.Selection, n int) *goquery.Selection {
	return data.Slice(n, n+1)
}
func timesplit(data *goquery.Selection) []string {
	ret := []string{}
	d := strings.TrimSpace(data.Text())
	ret = strings.Fields(d)
	return ret
}
