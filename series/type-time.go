package series

import (
	"fmt"
	"github.com/golang-module/carbon"
	"github.com/spf13/cast"
	"math"
)

type timeElement struct {
	e   carbon.Carbon
	nan bool
}

func (e *timeElement) Set(value interface{}) {
	e.nan = false
	switch value.(type) {
	case string:
		if value.(string) == "NaN" {
			e.nan = true
			return
		}
		e.e = carbon.Parse(value.(string))
	case int, int8, int16, int32, int64:
		n := cast.ToInt64(value)
		e.e = carbon.CreateFromTimestamp(n)
	case float64, float32:
		//f := value.(float64)
		f := cast.ToFloat64(value)
		if math.IsNaN(f) ||
			math.IsInf(f, 0) ||
			math.IsInf(f, 1) {
			e.nan = true
			return
		}
		e.e = carbon.CreateFromTimestamp(int64(f))
	case bool:
		e.nan = true
		return
	case Element:
		e.e = value.(Element).Time()
		//v, err := value.(Element).Int()
		//if err != nil {
		//	e.nan = true
		//	return
		//}
		//e.e = int64(v)
	default:
		e.nan = true
		return
	}
	return
}

func (e timeElement) Copy() Element {
	if e.IsNA() {
		return &timeElement{carbon.Parse(""), true}
	}
	return &timeElement{e.e, false}
}

func (e timeElement) IsNA() bool {
	if e.nan {
		return true
	}
	return false
}

func (e timeElement) Type() Type {
	return Time
}

func (e timeElement) Val() ElementValue {
	if e.IsNA() {
		return nil
	}
	return e.String()
}

func (e timeElement) String() string {
	if e.IsNA() {
		return "NaN"
	}
	return e.e.ToDateTimeString()
}

func (e timeElement) Int() (int, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to int")
	}
	return int(e.e.ToTimestampWithMicrosecond()), nil
}

func (e timeElement) Float() float64 {
	if e.IsNA() {
		return math.NaN()
	}
	return float64(e.e.ToTimestampWithMicrosecond())
}

func (e timeElement) Bool() (bool, error) {
	if e.IsNA() {
		return false, fmt.Errorf("can't convert NaN to bool")
	}
	return !e.e.IsZero(), nil
}

func (e timeElement) Time() carbon.Carbon {
	return e.e
}

func (e timeElement) Eq(elem Element) bool {
	i, err := elem.Int()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e.ToTimestampWithMicrosecond() == int64(i)
}

func (e timeElement) Neq(elem Element) bool {
	i, err := elem.Int()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e.ToTimestampWithMicrosecond() != int64(i)
}

func (e timeElement) Less(elem Element) bool {
	i, err := elem.Int()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e.ToTimestampWithMicrosecond() < int64(i)
}

func (e timeElement) LessEq(elem Element) bool {
	i, err := elem.Int()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e.ToTimestampWithMicrosecond() <= int64(i)
}

func (e timeElement) Greater(elem Element) bool {
	i, err := elem.Int()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e.ToTimestampWithMicrosecond() > int64(i)
}

func (e timeElement) GreaterEq(elem Element) bool {
	i, err := elem.Int()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e.ToTimestampWithMicrosecond() >= int64(i)
}
