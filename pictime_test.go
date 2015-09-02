// (C) Philip Schlump, 2014.

package pictime

import (
	"testing"
	"time"
)
	// "fmt"

type TestCase struct {
	format, value string
}

/*
	"Monday": "Monday",     // Weekday name
	"Month"	: "January",    // Month name
	"YYYY"	: "2006",       // Year 4 digit
	"yyyy"	: "2006",       // Year 2 digit
	"ddd"	: "Mon",        // 3 char weekday name
	"MMM"	: "Jan",        // 3 char month name
	"Mon"	: "Jan",        // 3 char month name
	"am"	: "pm",         // AM or PM
	"AM"	: "PM",         // AM or PM
	"DD"	: "02",         // Day of the month e.g. [01,31]
	"HH"	: "15",         // Hour (24-hour clock) e.g. [00,23]
	"hh"	: "3",          // Hour (12-hour clock) e.g. [01,12]
	"MM"	: "01",         // Month e.g. [01,12]
	"mm"	: "04",         // Minute e.g. [00,59]
	"pm"	: "pm",         // AM or PM
	"PM"	: "PM",         // AM or PM
	"RR"	: "06",         // Year 
	"ss"	: "05",         // Second as e.g. [00,61]
	"tt"	: "pm",         // AM or PM
	"TT"	: "PM",         // AM or PM
	"TZ"	: "MST",        // Time zone 
	"yy"	: "06",         // Year 
	"YY"	: "06",         // Year 
	"."		: ".",          // All other characters
*/

var testTime = time.Date(2014, time.January, 10, 23, 31, 32, 93, time.UTC)
var testCases = []*TestCase{
	&TestCase{"ddd", "Fri"},
	&TestCase{"Monday", "Friday"},
	&TestCase{"MMM", "Jan"},
	&TestCase{"Month", "January"},
	&TestCase{"DD", "10"},
	&TestCase{"HH", "23"},
	&TestCase{"MM", "01"},
	&TestCase{"mm", "31"},
	&TestCase{"PM", "PM"},
	&TestCase{"ss", "32"},
	&TestCase{"nnn", "930"},
	&TestCase{"yy", "14"},
	&TestCase{"YY", "14"},
	&TestCase{"yyyy", "2014"},
	&TestCase{"YYYY", "2014"},
	&TestCase{"TZ", "UTC"},

	// Escape
	&TestCase{"YYYY/MM/DD hh:mm:ss", "2014/01/10 11:31:32"},
	&TestCase{"YYYY-MM-DD hh:mm:ss", "2014-01-10 11:31:32"},
	// In a string
	&TestCase{"/aaaa/YYYY/mm/bbbb", "/aaaa/2014/31/bbbb"},
	// No Format - get rid of value
	&TestCase{"", ""},
}

func TestPictureFormats(t *testing.T) {
	for _, v := range testCases {
		result, err := Format(v.format, testTime)
		if err != nil {
			t.Fatalf("Error format=[%s] error=%v", v.format, err)
		}
		if result != v.value {
			t.Fatalf("Error format=[%s]: results=[%s] expected=[%s]", v.format, result, v.value)
		}
	}
}

