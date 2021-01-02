package util

import (
	"fmt"
	"strings"
	"time"
)

var replacer = strings.NewReplacer("*", "\\*", "_", "\\_")

func EscapeString(message string) string {
	return replacer.Replace(message)
}

func FormatTime(timeData time.Time, location string) (string, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return "", err
	}
	t := timeData.In(loc)
	zone, _ := t.Zone()
	return fmt.Sprintf("%s, %02d %s %d at %02d:%02d %s", t.Weekday().String(), t.Day(), t.Month().String(), t.Year(), t.Hour(), t.Minute(), zone), nil
}
