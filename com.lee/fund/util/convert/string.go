package convert

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func ChangeType(value string, t reflect.Type) interface{} {
	switch t.Kind() {
	case reflect.String:
		return ToString(value)
	case reflect.Int32:
		return StringToInt32(value, 0)
	case reflect.Int64:
		return StringToInt64(value, 0)
	case reflect.Float32:
		return StringToFloat32(value, 0)
	case reflect.Float64:
		return StringToFloat64(value, 0)
	case reflect.Bool:
		return StringToBool(value, false)
	}
	//TODO:执行到这里时会出错，待重构
	return reflect.ValueOf(value).Convert(t)
}

/********** ToString **********/
func ToString(obj interface{}) string {
	switch v := obj.(type) {
	case string:
		return v
	case *string:
		return *v
	case []byte:
		return string(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case int:
		return strconv.Itoa(v)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	}

	return fmt.Sprint(obj)
}

func BoolToString(b bool) string {
	return strconv.FormatBool(b)
}

func Int8ToString(i int8) string {
	return strconv.FormatInt(int64(i), 10)
}

func Int16ToString(i int16) string {
	return strconv.FormatInt(int64(i), 10)
}

func Int32ToString(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func Float32ToString(f float32) string {
	return strconv.FormatFloat(float64(f), 'f', -1, 32)
}

func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func SliceToString(sep string, args ...interface{}) string {
	if args == nil || len(args) <= 0 {
		return ""
	}

	arr := make([]string, len(args))
	for i, arg := range args {
		arr[i] = ToString(arg)
	}

	return strings.Join(arr, sep)
}

func Int32SliceToString(sep string, args ...int32) string {
	if args == nil || len(args) <= 0 {
		return ""
	}

	arr := make([]string, len(args))
	for i, arg := range args {
		arr[i] = Int32ToString(arg)
	}

	return strings.Join(arr, sep)
}

/********** StringTo **********/

// valid string: 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False
func StringToBool(s string, defaultValue bool) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return defaultValue
	}

	return b
}

func StringToInt8(s string, defaultValue int8) int8 {
	i, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		return defaultValue
	}

	return int8(i)
}

func StringToInt16(s string, defaultValue int16) int16 {
	i, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		return defaultValue
	}

	return int16(i)
}

func StringToInt32(s string, defaultValue int32) int32 {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return defaultValue
	}

	return int32(i)
}

func StringToInt64(s string, defaultValue int64) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return defaultValue
	}

	return i
}

func StringToInt(s string, defaultValue int) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return defaultValue
	}

	return int(i)
}

func StringToFloat32(s string, defaultValue float32) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return defaultValue
	}

	return float32(f)
}

func StringToFloat64(s string, defaultValue float64) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return defaultValue
	}

	return f
}

func ToMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
