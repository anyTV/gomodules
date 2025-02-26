package ftime

import "time"

const (
	Zero        = time.Second * 0
	Hour        = time.Hour
	Second      = time.Second
	Minute      = time.Minute
	Day         = 24 * time.Hour
	FiveMinutes = 5 * Minute
	TenMinutes  = 10 * Minute
	TenKSeconds = 10 * 1000 * time.Second
	OneHour     = 1 * time.Hour
	TenHours    = 10 * time.Hour
	ThirtyDays  = 30 * Day
)
