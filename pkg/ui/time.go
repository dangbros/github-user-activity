package ui

import (
	"fmt"
	"time"
)

func RelativeTime(isoTime string) string {
	t, err := time.Parse(time.RFC3339, isoTime)
	if err != nil {
		return ""
	}

	d := time.Since(t)

	switch {
	case d < time.Minute:
		return "just now"
	case d < time.Hour:
		return plural(int(d.Minutes()), "m")
	case d < 24*time.Hour:
		return plural(int(d.Hours()), "h")
	case d < 7*24*time.Hour:
		return plural(int(d.Hours()/24), "d")
	default:
		return plural(int(d.Hours()/(7*24)), "w")
	}

}

func plural(value int, unit string) string {
	return fmt.Sprintf("%d%s ago", value, unit)
}
