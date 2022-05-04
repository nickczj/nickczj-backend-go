package service

import "time"

func CurrentTray() int {
	loc, _ := time.LoadLocation("Asia/Singapore")
	firstTray := time.Date(2021, time.November, 2, 0, 0, 0, 0, loc)
	timeSince := time.Now().Sub(firstTray).Hours()
	return int(timeSince/(24*14)) + 1
}
