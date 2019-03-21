package config

import (
	"com.lee/fund/util/convert"
	"strconv"
	"strings"
	"time"
)

type SettingMap map[string]string

func (m SettingMap) String(key string, defaultValue string) string {
	v, ok := m[key]
	if ok {
		return v
	}

	return defaultValue
}

func (m SettingMap) Int(key string, defaultValue int) int {
	v, ok := m[key]
	if ok {
		i, err := strconv.Atoi(v)
		if err == nil {
			return i
		}
	}

	return defaultValue
}

func (m SettingMap) Int32(key string, defaultValue int32) int32 {
	v, ok := m[key]
	if ok {
		i, err := strconv.ParseInt(v, 10, 32)
		if err == nil {
			return int32(i)
		}
	}

	return defaultValue
}

func (m SettingMap) Bool(key string, defaultValue bool) bool {
	v, ok := m[key]
	if ok {
		b, err := strconv.ParseBool(v)
		if err == nil {
			return b
		}
	}

	return defaultValue
}

func (m SettingMap) Time(key string, defaultValue time.Time) time.Time {
	v, ok := m[key]
	if ok {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", v, time.Local)
		if err != nil {
			t, err = time.ParseInLocation("2006-01-02", v, time.Local)
		}
		if err == nil {
			return t
		}
	}

	return defaultValue
}

func (m SettingMap) StringSlice(key string) []string {
	v, ok := m[key]
	if !ok {
		return nil
	}

	return strings.Split(v, ",")
}

func (m SettingMap) Int32Slice(key string) []int32 {
	v, ok := m[key]
	if !ok {
		return nil
	}

	strs := strings.Split(v, ",")
	array := make([]int32, len(strs))
	for i := 0; i < len(strs); i++ {
		array[i] = convert.StringToInt32(strs[i], 0)
	}

	return array
}
