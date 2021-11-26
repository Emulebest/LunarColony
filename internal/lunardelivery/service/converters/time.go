package converters

import (
	"fmt"
	"time"
)

const (
	// Layout is an example date that shows the layout for time format
	Layout = "2006-01-02 15:04:05 MST"
	// LunarZeroDate is a start date of LST according to the UTC
	LunarZeroDate = "1969-07-21 02:56:15 UTC"
	// ConversionPoint 30 seconds on the Moon = 'ConversionPoint' seconds on Earth
	ConversionPoint = 29.530589
)

// StringToUTC converts string to the UTC format or returns an error
func StringToUTC(utc string) (time.Time, error) {
	parsed, err := time.Parse(Layout, utc)
	if err != nil {
		return time.Time{}, nil
	}
	return parsed, nil
}

// LunarElapsedSeconds checks the amount of seconds that have elapsed from the first Lunar Expedition
// Neil Armstrong set foot on the Moon surface on July 21th 1969 at 02:56:15 UT
func LunarElapsedSeconds(earthTime time.Time) float64 {
	elapsed := earthTime.Sub(GetLunarZeroDate())
	return elapsed.Seconds()
}

// ConvertEarthToMoonSec converts seconds from Earth to Moon ones
// 30 seconds on the Moon = 29.530589 seconds on Earth
func ConvertEarthToMoonSec(earthSeconds float64) (moonSeconds int64) {
	return int64(earthSeconds * 30 / ConversionPoint)
}

// FormatLST convert the amount of seconds according to the Lunar Standard Time format
func FormatLST(moonSeconds int64) string {
	moonMinutes, moonSeconds := divMod(moonSeconds, 60)
	moonHours, moonMinutes := divMod(moonMinutes, 60)
	moonCycles, moonHours := divMod(moonHours, 24)
	moonDays, moonCycles := divMod(moonCycles, 30)
	if moonCycles+1 > 30 {
		moonCycles = 1
		moonDays++
	} else {
		moonCycles++
	}
	moonYears, moonDays := divMod(moonDays, 12)
	if moonDays+1 > 12 {
		moonDays = 1
		moonYears++
	} else {
		moonDays++
	}
	moonYears++
	// We need to add +1 to years, days, cycles because time on Moon started from Year 1, Day 1, Cycle 1
	return fmt.Sprintf("%s-%s-%s delta %s:%s:%s", addTrailingZero(moonYears), addTrailingZero(moonDays),
		addTrailingZero(moonCycles), addTrailingZero(moonHours),
		addTrailingZero(moonMinutes), addTrailingZero(moonSeconds))
}

// GetLunarZeroDate gives you the lunar start time in Golang format
func GetLunarZeroDate() time.Time {
	date, err := time.Parse(Layout, LunarZeroDate)
	if err != nil {
		panic(err)
	}
	return date
}

// divMod is a utility function that given numerator and denominator, gives you quotient and remainder
func divMod(numerator int64, denominator int64) (quotient int64, remainder int64) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}

// addTrailingZero adds a trailing zero to the number if it is less than 10
func addTrailingZero(number int64) string {
	if number < 10 {
		return fmt.Sprintf("0%d", number)
	} else {
		return fmt.Sprintf("%d", number)
	}
}
