package id

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sony/sonyflake"

	"snowflake-id-generator/utils/timeutil"
)

func ExtractStartAndEndTimeFromID(idStr string) (time.Time, time.Time, error) {
	idTime, err := ExtractTimeFromID(idStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return timeutil.StartOfDay(idTime), timeutil.EndOfDay(idTime), nil
}

func ExtractTimeFromID(idStr string) (time.Time, error) {
	_, time, err := ExtractID(idStr)
	return time, err
}

func ExtractID(idStr string) (uint64, time.Time, error) {
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, time.Time{}, fmt.Errorf("invalid SnowflakeID format")
	}

	elapsedTime := sonyflake.ElapsedTime(id)
	currentTime := DefaultStartTime.Add(elapsedTime)

	return id, currentTime, nil
}
