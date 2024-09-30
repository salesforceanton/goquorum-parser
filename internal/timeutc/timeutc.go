package timeutc

import "time"

func Now() time.Time {
	return time.Now().UTC()
}

func Unix(date int64) time.Time {
	return time.Unix(date, 0).UTC()
}
