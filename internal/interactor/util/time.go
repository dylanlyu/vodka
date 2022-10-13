package util

import "time"

func ChangeToUTC(lasting time.Time) time.Time {
	return lasting.UTC()
}

func ChangeToRFC3339(lasting time.Time) string {
	return lasting.Format(time.RFC3339)
}

func NowToUTC() time.Time {
	return time.Now().UTC()
}

func UTCToRFC3339() string {
	utc := NowToUTC()
	return ChangeToRFC3339(utc)
}

func GetAge(birthday time.Time) (age int) {
	if birthday.IsZero() {
		return 0
	}

	now := NowToUTC()
	age = now.Year() - birthday.Year()
	if int(now.Month()) < int(birthday.Month()) || int(now.Day()) < int(birthday.Day()) {
		age--
	}
	return age
}
