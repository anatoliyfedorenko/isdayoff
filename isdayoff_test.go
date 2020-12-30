package isdayoff

import (
	"testing"
	"time"
)

func TestIsLeap(t *testing.T) {
	client := New()
	leap, err := client.IsLeap(2020)
	if err != nil {
		t.Error(err)
	}
	if leap != true {
		t.Errorf("should be true, equal: %v", leap)
	}
}

func TestGetByYear(t *testing.T) {
	client := New()
	days, err := client.GetBy(Params{Year: 2020})
	if err != nil {
		t.Error(err)
	}
	if len(days) != 366 {
		t.Errorf("should be 366, equal: %v", len(days))
	}
}

func TestGetByDay(t *testing.T) {
	client := New()
	month := time.January
	day := 1
	countryCode := CountryCodeKazakhstan
	pre := false
	covid := false
	days, err := client.GetBy(Params{
		Year:        2020,
		Month:       &month,
		Day:         &day,
		CountryCode: &countryCode,
		Pre:         &pre,
		Covid:       &covid,
	})
	if err != nil {
		t.Error(err)
	}
	if len(days) != 1 {
		t.Errorf("should be 1, equal: %v", len(days))
	}
}
