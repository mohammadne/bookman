package utils

import "time"

const (
	apiLayout      = "2006-01-02T15:04:05Z"
	databaseLayout = "2006-01-02 15:04:05"
)

func Now() time.Time {
	return time.Now().UTC()
}

func NowString() string {
	return Now().Format(apiLayout)
}

func NowDatabseFormatString() string {
	return Now().Format(databaseLayout)
}
