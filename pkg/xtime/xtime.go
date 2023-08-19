package xtime

import (
	"context"
	"database/sql/driver"
	"strconv"
	"time"
)

// Time be used to MySql timestamp converting.
type Time int64

// Scan scan time.
func (t *Time) Scan(src any) (err error) {
	switch sc := src.(type) {
	case time.Time:
		if sc.IsZero() {
			return
		}
		*t = Time(sc.Unix())
	case string:
		var i int64
		i, err = strconv.ParseInt(sc, 10, 64)
		*t = Time(i)
	}
	return
}

// Value get time value.
func (t Time) Value() (driver.Value, error) {
	return time.Unix(int64(t), 0), nil
}

// Time get time.
func (t Time) Time() time.Time {
	return time.Unix(int64(t), 0)
}

func (t *Time) FromDB(bs []byte) error {
	timeStr := string(bs)
	ti, err := time.ParseInLocation("2006-01-02T15:04:05", timeStr[:19], time.Local)
	if err != nil {
		return err
	}
	if ti.IsZero() {
		return nil
	}
	*t = Time(ti.Unix())
	return nil
}

func (t Time) ToDB() ([]byte, error) {
	unix := time.Unix(int64(t), 0)
	return []byte(unix.String()), nil
}

// Duration be used json unmarshal string time, like 1s, 500ms.
type Duration time.Duration

// UnmarshalText unmarshal text to duration.
func (d *Duration) UnmarshalText(text []byte) error {
	tmp, err := time.ParseDuration(string(text))
	if err == nil {
		*d = Duration(tmp)
	}
	return err
}

// UnitTime duration parse to unit, such as "300ms", "1h30m" or "2h10s".
func (d *Duration) UnitTime() string {
	return DurationToUnit(time.Duration(*d))
}

// Shrink will decrease the duration by comparing with context's timeout duration and return new timeout\context\CancelFunc.
func (d Duration) Shrink(c context.Context) (Duration, context.Context, context.CancelFunc) {
	if deadline, ok := c.Deadline(); ok {
		if ctimeout := time.Until(deadline); ctimeout < time.Duration(d) {
			// deliver small timeout
			return Duration(ctimeout), c, func() {}
		}
	}
	ctx, cancel := context.WithTimeout(c, time.Duration(d))
	return d, ctx, cancel
}
