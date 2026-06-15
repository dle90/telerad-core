package types

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

const DateLayout = "02/01/2006" // matches Go's reference time format for day/month/year

// Date maps to PostgreSQL `date` type — stores year/month/day without timezone.
// JSON form is always "DD/MM/YYYY".
type Date time.Time

func (d Date) Time() time.Time {
	return time.Time(d)
}

func (d Date) String() string {
	return time.Time(d).Format(DateLayout)
}

func (d Date) IsZero() bool {
	return time.Time(d).IsZero()
}

func (d Date) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + d.String() + `"`), nil
}

func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" || s == "null" {
		*d = Date{}
		return nil
	}
	t, err := time.Parse(DateLayout, s)
	if err != nil {
		return fmt.Errorf("invalid date %q, expected %s: %w", s, DateLayout, err)
	}
	*d = Date(t)
	return nil
}

func (d *Date) Scan(value any) error {
	if value == nil {
		*d = Date{}
		return nil
	}
	var s string
	switch v := value.(type) {
	case time.Time:
		*d = Date(v)
		return nil
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		return fmt.Errorf("Date.Scan: cannot scan %T into Date", value)
	}
	for _, layout := range []string{"2006-01-02", DateLayout, time.RFC3339} {
		if t, err := time.Parse(layout, s); err == nil {
			*d = Date(t)
			return nil
		}
	}
	return fmt.Errorf("Date.Scan: invalid date string %q", s)
}

func (d Date) Value() (driver.Value, error) {
	if d.IsZero() {
		return nil, nil
	}
	return time.Time(d).Format("2006-01-02"), nil
}
