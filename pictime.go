package pictime

/*

Implementation of Picture date/time format in Go
See http://www.codeproject.com/Articles/576178/cast-convert-format-try-parse-date-and-time-sql#2_2

Example:
	str, err := pictime.Format("YYYY/mm/dd", time.Now()) // 2014/05/07

Copyright (C) Philip Schlump, 2014-2016.

*/

import (
	"fmt"
	"regexp"
	"time"
)

type pictimeFmt struct {
	Format     string
	TimeFormat string
}

var setup = []pictimeFmt{
	{"Monday", "Monday"}, // Weekday name
	{"Month", "January"}, // Month name
	{"YYYY", "2006"},     // Year 4 digit
	{"yyyy", "2006"},     // Year 2 digit
	{"ddd", "Mon"},       // 3 char weekday name
	{"MMM", "Jan"},       // 3 char month name
	{"Mon", "Jan"},       // 3 char month name
	{"YYY", "006"},       // 3 char year
	{"yyy", "006"},       // 3 char year
	{"nnn", "000"},       // 3 char fractional seconds
	{"am", "pm"},         // AM or PM
	{"AM", "PM"},         // AM or PM
	{"DD", "02"},         // Day of the month e.g. [01,31]
	{"HH", "15"},         // Hour (24-hour clock) e.g. [00,23]
	{"hh", "3"},          // Hour (12-hour clock) e.g. [01,12]
	{"MM", "01"},         // Month e.g. [01,12]
	{"mm", "04"},         // Minute e.g. [00,59]
	{"pm", "pm"},         // AM or PM
	{"PM", "PM"},         // AM or PM
	{"RR", "06"},         // Year
	{"ss", "05"},         // Second as e.g. [00,61]
	{"tt", "pm"},         // AM or PM
	{"TT", "PM"},         // AM or PM
	{"TZ", "MST"},        // Time zone
	{"yy", "06"},         // Year
	{"YY", "06"},         // Year
	{".", "."},           // All other characters
}

var conv map[string]string

var fmtRe *regexp.Regexp

func init() {
	conv = make(map[string]string, len(setup))
	s := "("
	com := ""
	for _, v := range setup {
		s = s + com + "(" + v.Format + ")"
		com = "|"
		conv[v.Format] = v.TimeFormat
	}
	s += ")"
	// fmt.Printf ( "RE=[%s]\n", s )
	fmtRe = regexp.MustCompile(s)
}

func repl(match string, t time.Time) string {

	if format, ok := conv[match]; ok {
		// fmt.Printf ( "ms=%v", testTime.Nanosecond() )
		if format == "." {
			return match
		} else if match == "nnn" {
			return ZeroPadRight(3, fmt.Sprintf("%v", t.Nanosecond()))
		} else {
			return t.Format(format)
		}
	} else {
		return match
	}
	// panic(fmt.Errorf("unknown picture format directive - %s", match))
}

// Format applies the  picture format to the time 't' and returns the string or an error
func Format(format string, t time.Time) (result string, err error) {
	defer func() {
		if e := recover(); e != nil {
			result = ""
			err = e.(error)
		}
	}()

	fn := func(match string) string {
		return repl(match, t)
	}
	return fmtRe.ReplaceAllStringFunc(format, fn), nil
}

// ZeroPadRight padds an string 's' with 0's on the right to the desired length, 'l'
func ZeroPadRight(l int, s string) string {
	return PadStrRight(l, "0", s)
}

// PadStrRight padds string on the right.  to a length of 'l' with 'w'
func PadStrRight(l int, w string, s string) string {
	if len(s) >= l {
		return s
	}
	k := l - len(s)
	t := ""
	for i := 0; i < k; i++ {
		t += w
	}
	return s + t
}
