package utils

import "time"

func GetTimeNowInUTC() string {
	return time.Now().UTC().Format(time.RFC3339)
}
