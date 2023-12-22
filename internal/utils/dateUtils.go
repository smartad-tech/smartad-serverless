package utils

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func UnixTimestampToDateString(unixTimestamp string) string {
	intTimestamp, err := strconv.ParseInt(unixTimestamp, 10, 64)
	if err != nil {
		log.Fatalf("Error during parsing timestamp from ViewEntity. Error: %s", err.Error())
	}
	t := time.Unix(intTimestamp, 0)
	return fmt.Sprintf("%02d.%02d.%04d", t.Day(), t.Month(), t.Year())
}

type SmartDate struct {
	date time.Time
}

// NewSmartDateFromString input date format is DD.MM.YY
func NewSmartDateFromString(stringDate string) (SmartDate, error) {
	parsedDate, err := time.Parse("02.01.2006", stringDate)
	if err != nil {
		return SmartDate{}, nil
	}
	return SmartDate{date: parsedDate}, nil
}

func NewSmartDate() SmartDate {
	return SmartDate{date: time.Now()}
}

func NewSmartDateFromUnix(unixDate string) (SmartDate, error) {
	unixFormatTime, err := strconv.ParseInt(unixDate, 10, 64)
	if err != nil {
		return SmartDate{}, err
	}

	return SmartDate{date: time.Unix(unixFormatTime, 0)}, nil
}

func (d SmartDate) Unix() int64 {
	return d.date.Unix()
}

func (d SmartDate) ToString() string {
	return fmt.Sprintf("%02d.%02d.%04d", d.date.Day(), d.date.Month(), d.date.Year())
}
