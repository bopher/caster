package caster

import (
	"fmt"
	"strconv"
	"strings"
)

func asBool(v interface{}) (bool, bool) {
	b, e := strconv.ParseBool(fmt.Sprint(v))
	return b, e == nil
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
