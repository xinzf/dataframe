package series

import (
	"fmt"
	"github.com/golang-module/carbon"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cast"
	"math"
	"reflect"
	"strings"
)

type arrayElement struct {
	//e   Series
	e        []interface{}
	dataType string
	nan      bool
	//t Type
}

func (e *arrayElement) Set(value interface{}) {
	e.e = make([]interface{}, 0)
	e.dataType = "string"

	detectDataType := func(v interface{}) string {
		switch v.(type) {
		case string:
			return "string"
		case int, int8, int32, int16, int64, uint, uint8, uint16, uint32, uint64:
			return "integer"
		case float64, float32:
			return "float"
		case bool:
			return "boolean"
		case Element:
			return "element"
		default:
			return "string"
		}
	}

	switch reflect.Indirect(reflect.ValueOf(value)).Kind() {
	case reflect.Slice:
		byts, err := jsoniter.Marshal(value)
		if err != nil {
			e.nan = true
			return
		}
		if err = jsoniter.Unmarshal(byts, &e.e); err != nil {
			e.nan = true
			return
		}

		if len(e.e) == 0 {
			return
		}

	case reflect.Map, reflect.Struct, reflect.Ptr:

	default:
		e.e = append(e.e, value)

	}
	e.dataType = detectDataType(e.e[0])
	e.nan = false
}

func (e arrayElement) Eq(element Element) bool {
	if e.Type() != element.Type() {
		return false
	}

	targetValue := strings.Split(element.String(), ",")
	currValue := strings.Split(e.String(), ",")
	if len(currValue) != len(targetValue) {
		return false
	}

	finder := func(v string) bool {
		for _, vv := range targetValue {
			if v == vv {
				return true
			}
		}
		return false
	}

	match := 0
	for _, v := range currValue {
		if finder(v) {
			match++
		}
	}

	if match == len(e.e) {
		return true
	}
	return false
}

func (e arrayElement) Neq(element Element) bool {
	return !e.Eq(element)
	//if e.IsNA() || element.IsNA() {
	//	return false
	//}

	//for i := 0; i < e.e.elements.Len(); i++ {
	//	if e.e.elements.Elem(i).Eq(element) == false {
	//		return true
	//	}
	//}
	//return false
}

func (a arrayElement) Less(element Element) bool {
	return false
}

func (a arrayElement) LessEq(element Element) bool {
	return false
}

func (a arrayElement) Greater(element Element) bool {
	return false
}

func (a arrayElement) GreaterEq(element Element) bool {
	return false
}

func (e arrayElement) Copy() Element {
	if e.IsNA() {
		return &arrayElement{
			e:   []interface{}{},
			nan: true,
		}
	}

	return &arrayElement{
		e:   e.e,
		nan: e.nan,
	}
}

func (e arrayElement) Val() ElementValue {
	if e.nan {
		return make([]interface{}, 0)
	}

	switch e.dataType {
	//case "string":
	//	strs := make([]string, 0)
	//	for _, v := range e.e {
	//		ele := new(stringElement)
	//		ele.Set(v)
	//		strs = append(strs, ele.String())
	//	}
	//	return strs
	case "integer":
		ints := make([]int, 0)
		for _, v := range e.e {
			ele := new(intElement)
			ele.Set(v)
			num, _ := ele.Int()
			ints = append(ints, num)
		}
		return ints
	case "float":
		floats := make([]float64, 0)
		for _, v := range e.e {
			ele := new(floatElement)
			ele.Set(v)
			floats = append(floats, ele.Float())
		}
		return floats
	case "boolean":
		bools := make([]bool, 0)
		for _, v := range e.e {
			ele := new(boolElement)
			ele.Set(v)
			bool, _ := ele.Bool()
			bools = append(bools, bool)
		}
	default:
		return []interface{}{}
	}
	return []interface{}{}
}

func (e arrayElement) String() string {
	vals := make([]string, 0)
	for _, v := range e.e {
		val, err := cast.ToStringE(v)
		if err == nil {
			vals = append(vals, val)
		}
	}
	return strings.Join(vals, ",")
}

func (a arrayElement) Int() (int, error) {
	return 0, fmt.Errorf("can't convert array to int")
}

func (a arrayElement) Float() float64 {
	return math.NaN()
}

func (a arrayElement) Bool() (bool, error) {
	return false, fmt.Errorf("can't convert array to bool")
}

func (a arrayElement) Time() carbon.Carbon {
	return carbon.Carbon{}
}

func (e arrayElement) IsNA() bool {
	return e.nan
	//if e.nan {
	//	return true
	//}
	//for _, b := range e.e.IsNaN() {
	//	if b {
	//		return false
	//	}
	//}
	//return true
}

func (a arrayElement) Type() Type {
	return Array
}
