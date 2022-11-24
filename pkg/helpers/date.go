package helpers

import "time"

func DateFormatEN() string {
	currentTime := time.Now()
	dateEN := currentTime.Format("2006-01-02")
	return dateEN
}
