package enum

/*
type Weekday int

const (
  Sunday    Weekday = 1
  Monday    Weekday = 2
  Tuesday   Weekday = 3
  Wednesday Weekday = 4
  Thursday  Weekday = 5
  Friday    Weekday = 6
  Saturday  Weekday = 7
)

var names := [...]string{
  "Sunday",
  "Monday",
  "Tuesday",
  "Wednesday",
  "Thursday",
  "Friday",
  "Saturday"}

// WeekdayText - Returns the description of the weekday.
//fmt.Println(Sunday)    // prints 1
func (day Weekday) WeekdayText() string {
  if day < Sunday || day > Saturday {
    return "Unknown"
  }
  return names[day]
}

func (code int) WeekdayText() string {
  if code < Sunday || code > Saturday {
    return "Unknown"
  }
  return names[code]
}

// IsWeekend - Checks if the day is a weekend.
// fmt.Printf("Is Saturday a weekend day? %t\n", Saturday.IsWeekend()) // Is Saturday a weekend day? true
func (day Weekday) IsWeekend() bool {
  switch day {
  case Sunday, Saturday:
    return true
  default:
    return false
  }
}
*/