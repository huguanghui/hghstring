// Package hghstring formats time.Duration into a human readable format
package hghstring

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

var (
	units = []string{"years", "weeks", "days", "hours", "minutes", "seconds", "milliseconds"}
)

// Hghfmt holds the parsed duration and the original input duration
type Hghfmt struct {
	duration time.Duration
	input    string
	short    bool
}

// Parse creates a new *Hghfmt struct, return error if input is invalid.
func Parse(dinput time.Duration) *Hghfmt {
	input := dinput.String()
	return &Hghfmt{dinput, input, false}
}

// ParseShort creates a new *Hghfmt struct, short form, return error if input is invalid.
func ParseShort(dinput time.Duration) *Hghfmt {
	input := dinput.String()
	return &Hghfmt{dinput, input, true}
}

// ParseString creates a new *Hghfmt struct from a string.ParseString.
// returns an error if input is invalid.
func ParseString(input string) (*Hghfmt, error) {
	if input == "0" || input == "-0" {
		return nil, errors.New("hghfmt: missing unit in duration " + input)
	}
	duration, err := time.ParseDuration(input)
	if err != nil {
		return nil, err
	}
	return &Hghfmt{duration, input, false}, nil
}

// ParseStringShort creates a new *Hghfmt struct from a string, short form
// returns an error if input is invalid.
func ParseStringShort(input string) (*Hghfmt, error) {
	if input == "0" || input == "-0" {
		return nil, errors.New("hghfmt: missing unit in duration " + input)
	}
	duration, err := time.ParseDuration(input)
	if err != nil {
		return nil, err
	}
	return &Hghfmt{duration, input, true}, nil
}

func (d *Hghfmt) String() string {
	var duration string

	// Check for minus durations.
	if string(d.input[0]) == "-" {
		duration += "-"
		d.duration = -d.duration
	}

	// Convert duration
	seconds := int64(d.duration.Seconds()) % 60
	minutes := int64(d.duration.Minutes()) % 60
	hours := int64(d.duration.Hours()) % 24
	days := int64(d.duration/(24*time.Hour)) % 365 % 7

	// Edge case between 364 and 365 days.
	// We need to calculate weeks from what is left from years
	leftYearDays := int64(d.duration/(24*time.Hour)) % 365
	weeks := leftYearDays / 7

	if leftYearDays >= 364 && leftYearDays < 365 {
		weeks = 52
	}

	years := int64(d.duration/(24*time.Hour)) / 365
	milliseconds := int64(d.duration/time.Millisecond) -
		(seconds * 1000) - (minutes * 60000) - (hours * 3600000) -
		(days * 86400000) - (weeks * 604800000) - (years * 31536000000)

	// Create a map of the converted duration time.
	durationMap := map[string]int64{
		"milliseconds": milliseconds,
		"seconds":      seconds,
		"minutes":      minutes,
		"hours":        hours,
		"days":         days,
		"weeks":        weeks,
		"years":        years,
	}

	// Construct duration string.
	for _, u := range units {
		v := durationMap[u]
		strval := strconv.FormatInt(v, 10)
		switch {
		case v > 1:
			duration += strval + " " + u + " "
		case v == 1:
			duration += strval + " " + strings.TrimRight(u, "s") + " "
		case d.duration.String() == "0" || d.duration.String() == "0s":
			if strings.HasSuffix(d.input, string(u[0])) && !strings.Contains(d.input, "ms") {
				if u == "milliseconds" {
					continue
				}
				duration += strval + " " + u
			}
			if u == "milliseconds" {
				if strings.Contains(d.input, "ms") {
					duration += strval + " " + u
					break
				}
			}
			break
		case v == 0:
			continue
		}
	}
	duration = strings.TrimSpace(duration)

	if d.short {
		duration = strings.Join(strings.Split(duration, " ")[:2], " ")
	}

	return duration
}
