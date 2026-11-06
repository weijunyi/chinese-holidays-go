package holidays

import (
	"testing"
	"time"
)

func TestLoadData(t *testing.T) {
	events, err := loadData()
	if err != nil {
		t.Error(err)
	}

	if len(events) != 149 {
		t.Error(len(events))
	}
}

func TestBundled2027(t *testing.T) {
	queryer, err := BundleQueryer()
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		date    string
		holiday bool
	}{
		{"2027-01-01", true},  // 元旦 (Fri)
		{"2027-01-31", false}, // 春节调休上班 (Sun)
		{"2027-02-05", true},  // 春节 除夕 (Fri)
		{"2027-02-08", true},  // 春节 (Mon)
		{"2027-02-14", false}, // 春节调休上班 (Sun)
		{"2027-04-05", true},  // 清明节 (Mon)
		{"2027-05-05", true},  // 劳动节 (Wed)
		{"2027-05-08", false}, // 劳动节调休上班 (Sat)
		{"2027-06-09", true},  // 端午节 (Wed)
		{"2027-09-15", true},  // 中秋节 (Wed)
		{"2027-09-26", false}, // 国庆节调休上班 (Sun)
		{"2027-10-07", true},  // 国庆节 (Thu)
		{"2027-10-09", false}, // 国庆节调休上班 (Sat)
	}

	for _, c := range cases {
		d, err := time.ParseInLocation("2006-01-02", c.date, location)
		if err != nil {
			t.Fatal(err)
		}

		holiday, err := queryer.IsHoliday(d)
		if err != nil {
			t.Fatal(err)
		}
		if holiday != c.holiday {
			t.Errorf("%s: IsHoliday = %v, want %v", c.date, holiday, c.holiday)
		}

		workingday, err := queryer.IsWorkingday(d)
		if err != nil {
			t.Fatal(err)
		}
		if workingday == c.holiday {
			t.Errorf("%s: IsWorkingday = %v, want %v", c.date, workingday, !c.holiday)
		}
	}
}
