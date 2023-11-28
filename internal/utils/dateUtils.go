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
