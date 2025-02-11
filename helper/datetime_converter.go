package helper

import "time"

type DateTimeConverter struct {
}

func StringToDateTime(dateTime string) time.Time {
	parsed, _ := time.Parse("2006-01-02 15:04:05", dateTime)

	return parsed
}

func StringToDate(dateTime string) time.Time {
	parsed, _ := time.Parse("2006-01-02", dateTime)
	return parsed
}

func (d *DateTimeConverter) ConvertToDateString(datetime time.Time) string {
	return datetime.Format("2006-01-02")
}

func (d *DateTimeConverter) ConvertToDateTimeString(datetime time.Time) string {
	return datetime.Format("2006-01-02 15:04:05")
}

func ConvertToDateString(datetime time.Time) string {
	return datetime.Format("2006-01-02")
}

func ConvertToDateTimeString(datetime time.Time) string {
	return datetime.Format("2006-01-02 15:04:05")
}
