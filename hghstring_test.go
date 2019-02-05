package hghstring

import (
	"testing"
	"time"
)

var (
	testStrings []struct {
		test     string
		expected string
	}
	testTimes []struct {
		test     time.Duration
		expected string
	}
)

func TestParse(t *testing.T) {
	testTimes = []struct {
		test     time.Duration
		expected string
	}{
		{1 * time.Millisecond, "1 millisecond"},
		{1 * time.Second, "1 second"},
		{1 * time.Hour, "1 hour"},
		{1 * time.Minute, "1 minute"},
		{2 * time.Millisecond, "2 milliseconds"},
		{2 * time.Second, "2 seconds"},
		{2 * time.Minute, "2 minutes"},
		{1 * time.Hour, "1 hour"},
		{2 * time.Hour, "2 hours"},
		{10 * time.Hour, "10 hours"},
		{24 * time.Hour, "1 day"},
		{48 * time.Hour, "2 days"},
		{120 * time.Hour, "5 days"},
		{168 * time.Hour, "1 week"},
		{672 * time.Hour, "4 weeks"},
		{8759 * time.Hour, "52 weeks 23 hours"},
		{8760 * time.Hour, "1 year"},
		{17519 * time.Hour, "1 year 52 weeks 23 hours"},
		{17520 * time.Hour, "2 years"},
		{26279 * time.Hour, "2 years 52 weeks 23 hours"},
		{26280 * time.Hour, "3 years"},
		{201479 * time.Hour, "22 years 52 weeks 23 hours"},
		{201480 * time.Hour, "23 years"},
		{-1 * time.Second, "-1 second"},
		{-10 * time.Second, "-10 seconds"},
		{-100 * time.Second, "-1 minute 40 seconds"},
		{-1 * time.Millisecond, "-1 millisecond"},
		{-10 * time.Millisecond, "-10 milliseconds"},
		{-100 * time.Millisecond, "-100 milliseconds"},
	}

	for _, table := range testTimes {
		result := Parse(table.test).String()
		if result != table.expected {
			t.Errorf("Parse(%q).String() = %q. got %q, expected %q",
				table.test, result, result, table.expected)
		}
	}
}

func TestParseShort(t *testing.T) {
	testTimes = []struct {
		test     time.Duration
		expected string
	}{
		{1 * time.Millisecond, "1 millisecond"},
		{1 * time.Second, "1 second"},
		{1 * time.Hour, "1 hour"},
		{1 * time.Minute, "1 minute"},
		{2 * time.Millisecond, "2 milliseconds"},
		{2 * time.Second, "2 seconds"},
		{2 * time.Minute, "2 minutes"},
		{1 * time.Hour, "1 hour"},
		{2 * time.Hour, "2 hours"},
		{10 * time.Hour, "10 hours"},
		{24 * time.Hour, "1 day"},
		{48 * time.Hour, "2 days"},
		{120 * time.Hour, "5 days"},
		{168 * time.Hour, "1 week"},
		{672 * time.Hour, "4 weeks"},
		{8759 * time.Hour, "52 weeks"},
		{8760 * time.Hour, "1 year"},
		{17519 * time.Hour, "1 year"},
		{17520 * time.Hour, "2 years"},
		{26279 * time.Hour, "2 years"},
		{26280 * time.Hour, "3 years"},
		{201479 * time.Hour, "22 years"},
		{201480 * time.Hour, "23 years"},
		{-1 * time.Second, "-1 second"},
		{-10 * time.Second, "-10 seconds"},
		{-100 * time.Second, "-1 minute"},
		{-1 * time.Millisecond, "-1 millisecond"},
		{-10 * time.Millisecond, "-10 milliseconds"},
		{-100 * time.Millisecond, "-100 milliseconds"},
	}

	for _, table := range testTimes {
		result := ParseShort(table.test).String()
		if result != table.expected {
			t.Errorf("Parse(%q).String() = %q. got %q, expected %q",
				table.test, result, result, table.expected)
		}
	}
}

func TestParseString(t *testing.T) {
	testStrings = []struct {
		test     string
		expected string
	}{
		{"1ms", "1 millisecond"},
		{"2ms", "2 milliseconds"},
	}

	for _, table := range testStrings {
		d, err := ParseString(table.test)
		if err != nil {
			t.Errorf("%q", err)
		}
		result := d.String()
		if result != table.expected {
			t.Errorf("d.String() = %q. got %q, expected %q",
				table.test, result, table.expected)
		}
	}
}
