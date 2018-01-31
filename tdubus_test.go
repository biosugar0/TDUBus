package TDUBus

import (
	"fmt"
	"testing"
)

func TestGetTimeSchedyle(t *testing.T) {
	timetable := GetTimeSchedules()
	fmt.Println(timetable[2].Down.Takasaka["10"])
}
