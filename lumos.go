package main

import (
	"fmt"
	"strconv"
	"time"
)

var (
	systems   = []string{"BID", "BIB", "VISAP", "CASAP", "CONSOLIDATOR"}
	dateStart = time.Date(2015, time.September, 14, 0, 0, 0, 0, time.UTC)
)

func main() {
	duration := time.Since(dateStart)
	days := int(duration.Hours() / 24)
	weekth := days / 7
	if days%7 != 0 {
		weekth = days/7 + 1
	}
	position := weekth % len(systems)
	fmt.Println("days:" + strconv.Itoa(days) + "  weekth:" + strconv.Itoa(weekth))
	fmt.Println(systems[position-1])
}
