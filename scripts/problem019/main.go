package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Date(1901, time.January, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2000, time.December, 31, 0, 0, 0, 0, time.UTC)

	sundays := 0

	fmt.Printf("date range: %s to %s\n", start, end)

	for d := start; end.Compare(d) == 1; d = d.Add(86400000000000) {
		if d.Day() == 1 && d.Weekday() == time.Sunday {
			sundays += 1
		}
	}

	fmt.Println("sunday count:", sundays)
}
