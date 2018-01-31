# TDUBus

version:1.0

A package for getting bus timetable of TDU

## Function

### GetTimeSchedules()

* return value : []BusTimes

## Usage

### Monday ~ Friday

```
 timetable := TDUBus.GetTimeSchedules()
 fmt.Println(timetable[0])

// go to university from Takasaka
 fmt.Println(timetable[0].Up.Takasaka)

// go to Takasaka from university
 fmt.Println(timetable[0].Down.Takasaka)

// go to Takasaka from university (about ten o'clock)
 fmt.Println(timetable[0].Down.Takasaka["10"])

```

### Saturday

```
 timetable := TDUBus.GetTimeSchedules()
 fmt.Println(timetable[1])

// go to university from Takasaka
 fmt.Println(timetable[1].Up.Takasaka)

// go to Takasaka from university
 fmt.Println(timetable[1].Down.Takasaka)

// go to Takasaka from university (about ten o'clock)
 fmt.Println(timetable[1].Down.Takasaka["10"])

```

### Vacation (Monday ~ Friday)

```
 timetable := TDUBus.GetTimeSchedules()
 fmt.Println(timetable[2])

// go to university from Takasaka
 fmt.Println(timetable[2].Up.Takasaka)

// go to Takasaka from university
 fmt.Println(timetable[2].Down.Takasaka)

// go to Takasaka from university (about ten o'clock)
 fmt.Println(timetable[2].Down.Takasaka["10"])

```
