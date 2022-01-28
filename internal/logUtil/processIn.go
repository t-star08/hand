package logUtil

import (
	"fmt"
	"strconv"
)

func PublishID(n int) string {
	return fmt.Sprintf("%03d", n)
}

func PublishNextID(lastLog *LogUnit) string {
	lastNumID, _ := strconv.Atoi(lastLog.ID)
	return PublishID(lastNumID + 1)
}

func PublishPrevID(lastLog *LogUnit) string {
	lastNumID, _ := strconv.Atoi(lastLog.ID)
	return PublishID(lastNumID - 1)
}
