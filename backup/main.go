package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	matched, err := regexp.MatchString("\\d\\d\\d\\d-\\d\\d-\\d\\d/\\d\\d:\\d\\d:\\d\\d", "1997-01-21/00:00:00")
	fmt.Println(matched, err)
	const TimeFormat = "2006-01-02/15:04:05"
	t, err := time.Parse(TimeFormat, "2013-00-11/11:18:46")
	fmt.Println(t, err)
}
