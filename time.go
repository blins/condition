package condition

import (
	"errors"
	"time"
)

// проверяет на принадлежность времени ко временному интервалу
type TimeCondition struct {
	Start time.Time
	Stop  time.Time
}

func (condition *TimeCondition) ParseArgs(args []string) ([]string, error) {
	if len(args) == 0 {
		return nil, ErrExpectedValue("time")
	}
	var err error
	condition.Start, err = anyParseTime(args[0])
	if err != nil {
		return nil, err
	}
	condition.Stop = condition.Start
	shift := 1
	if len(args) > 1 {
		d, err := time.ParseDuration(args[1])
		if err == nil {
			condition.Stop = condition.Start.Add(d)
		}
		shift++
	}
	return args[shift:], nil
}

func (condition *TimeCondition) Check(value interface{}) bool {
	switch t := value.(type) {
	case time.Time:
		if t.Equal(condition.Start) || t.Equal(condition.Stop) {
			return true
		}
		if t.After(condition.Start) && t.Before(condition.Stop) {
			return true
		}
	}
	return false
}

func anyParseTime(str string) (time.Time, error) {
	formats := []string{time.ANSIC, time.Kitchen, time.RFC822, time.RFC822Z, time.RFC850, time.RFC1123,
		time.RFC1123Z, time.RFC3339, time.RFC3339Nano, time.RubyDate, time.StampMicro, time.StampMilli, time.StampNano,
		time.UnixDate, "2006-01-02T15:04:05", "2006-01-02T15:04"}
	for _, format := range formats {
		t, err := time.Parse(format, str)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, errors.New("not valid format")
}
