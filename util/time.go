package util

import (
	"time"
)

func UnixTimeToUTCDate(unixT int64) time.Time {
	return time.Unix(unixT, 0).UTC()
}
