package booking

import ("time"
        "fmt")

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
    t, _ := time.Parse(layout,date)
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
    t, _ := time.Parse(layout, date)
    result := time.Now().Compare(t)

    if result == 1 {
        return true
    } else {
        return false
    }
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
    t,_ := time.Parse(layout, date)
    hour, min, _ := t.Clock()
    current := fmt.Sprintf("%v:%v", hour, min)
	if current >= "12:00" && current < "18:00" {
        return true
    }
    return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
    t, _ := time.Parse(layout, date)
	//logic for hour and min
	hour, min, _ := t.Clock()
    current := fmt.Sprintf("%v:%v", hour, min)
    //logic for year month date and day
    year, month, d := t.Date()
    monthString :=  month.String()
    day := t.Weekday().String()
    desc := fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %s.", day, monthString, d, year, current)
    return desc
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	layout := "02 January, 2006 "
    date := "15 September, 2026"
    t, _ := time.Parse(layout, date)
    return t
}
