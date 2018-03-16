package TDUBus

import (
	"fmt"
	"testing"
)

func TestGetTimeSchedyle(t *testing.T) {
	timetable := GetTimeSchedules()
	fmt.Println(timetable[2].Down.Takasaka["10"])
}

func TestNext(t *testing.T) {
	B := Cli{}
	result, vacation := B.NextDown("takasaka")
	fmt.Println(result, vacation)
}
