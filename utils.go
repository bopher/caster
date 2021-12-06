package caster

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func asBool(v interface{}) (bool, bool) {
	s := strings.ToLower(fmt.Sprint(v))
	ok := s == "true" || s == "false"
	return s == "true", ok
}

func asInt64(v interface{}) (int64, bool) {
	i, e := strconv.ParseInt(fmt.Sprint(v), 10, 64)
	return i, e == nil
}

func asUint64(v interface{}) (uint64, bool) {
	u, e := strconv.ParseUint(fmt.Sprint(v), 10, 64)
	return u, e == nil
}

func asFloat32(v interface{}) (float32, bool) {
	f, e := strconv.ParseFloat(fmt.Sprint(v), 32)
	return float32(f), e == nil
}

func asFloat64(v interface{}) (float64, bool) {
	f, e := strconv.ParseFloat(fmt.Sprint(v), 64)
	return f, e == nil
}

func asString(v interface{}) (string, bool) {
	if reflect.TypeOf(v).Kind() == reflect.String {
		return reflect.ValueOf(v).String(), true
	}
	return "", false
}

func asSlice(v interface{}) []interface{} {
	res := make([]interface{}, 0)
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(v)
		for i := 0; i < s.Len(); i++ {
			if s.Index(i).CanInterface() {
				res = append(res, s.Index(i).Interface())
			}
		}
	}
	return res
}

func taggedError(tags []string, format string, args ...interface{}) error {
	_tags := ""
	for _, t := range tags {
		_tags = fmt.Sprintf("%s[%s] ", _tags, t)
	}
	return fmt.Errorf(_tags+format, args...)
}

func isErrorOf(tag string, err error) bool {
	return strings.Contains(err.Error(), "["+tag+"]")
}
