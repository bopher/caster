package caster

import (
	"math"
)

type Caster struct {
	data any
}

func (Caster) err(pattern string, params ...any) error {
	return taggedError([]string{"Caster"}, pattern, params...)
}

func (Caster) nilErr() error {
	return taggedError([]string{"Caster", "Caster.Nil"}, "value is nil!")
}

// IsNil check if value of caster is nil
func (c Caster) IsNil() bool {
	return c.data == nil
}

// Interface get raw value
func (c Caster) Interface() any {
	return c.data
}

// Slice parse data as []any
func (c Caster) Slice() []any {
	return asSlice(c.data)
}

// Bool parse data as boolean or return error on fail
func (c Caster) Bool() (bool, error) {
	if c.IsNil() {
		return false, c.nilErr()
	}

	if v, ok := asBool(c.data); ok {
		return v, nil
	}
	return false, c.err(`failed cast "%v" as bool!`, c.data)
}

// BoolSafe parse data as boolean or return fallback
func (c Caster) BoolSafe(fallback bool) bool {
	if v, err := c.Bool(); err == nil {
		return v
	}
	return fallback
}

// BoolSlice parse data as []bool or return fallback
func (c Caster) BoolSlice(fallback []bool) []bool {
	res := make([]bool, 0)
	for _, item := range c.Slice() {
		if v, ok := asBool(item); ok {
			res = append(res, v)
		}
	}

	if len(res) > 0 {
		return res
	}

	return fallback
}

// Int parse data as int or return error on fail
func (c Caster) Int() (int, error) {
	if c.IsNil() {
		return 0, c.nilErr()
	}

	if v, ok := asInt64(c.data); ok &&
		v >= math.MinInt &&
		v <= math.MaxInt {
		return int(v), nil
	}

	return 0, c.err("failed cast %v as int!", c.data)
}

// IntSafe parse data as int or return fallback
func (c Caster) IntSafe(fallback int) int {
	if v, err := c.Int(); err == nil {
		return v
	}
	return fallback
}

// IntSlice parse data as []int or return fallback
func (c Caster) IntSlice(fallback []int) []int {
	res := make([]int, 0)
	for _, item := range c.Slice() {
		if v, ok := asInt64(item); ok &&
			v >= math.MinInt &&
			v <= math.MaxInt {
			res = append(res, int(v))
		}
	}

	if len(res) > 0 {
		return res
	}

	return fallback
}

// Int8 parse data as int8 or return error on fail
func (c Caster) Int8() (int8, error) {
	if c.IsNil() {
		return 0, c.nilErr()
	}

	if v, ok := asInt64(c.data); ok &&
		v >= math.MinInt8 &&
		v <= math.MaxInt8 {
		return int8(v), nil
	}

	return 0, c.err("failed cast %v as int8!", c.data)
}

// Int8Safe parse data as int8 or return fallback
func (c Caster) Int8Safe(fallback int8) int8 {
	if v, err := c.Int8(); err == nil {
		return v
	}
	return fallback
}

// Int8Slice parse data as []int8 or return fallback
func (c Caster) Int8Slice(fallback []int8) []int8 {
	res := make([]int8, 0)
	for _, item := range c.Slice() {
		if v, ok := asInt64(item); ok &&
			v >= math.MinInt8 &&
			v <= math.MaxInt8 {
			res = append(res, int8(v))
		}
	}

	if len(res) > 0 {
		return res
	}

	return fallback
}

// Int16 parse data as int16 or return error on fail
func (c Caster) Int16() (int16, error) {
	if c.IsNil() {
		return 0, c.nilErr()
	}

	if v, ok := asInt64(c.data); ok &&
		v >= math.MinInt16 &&
		v <= math.MaxInt16 {
		return int16(v), nil
	}

	return 0, c.err("failed cast %v as int16!", c.data)
}

// Int16Safe parse data as int16  or return fallback
func (c Caster) Int16Safe(fallback int16) int16 {
	if v, err := c.Int16(); err == nil {
		return v
	}
	return fallback
}

// Int16Slice parse data as []int16 or return fallback
func (c Caster) Int16Slice(fallback []int16) []int16 {
	res := make([]int16, 0)
	for _, item := range c.Slice() {
		if v, ok := asInt64(item); ok &&
			v >= math.MinInt16 &&
			v <= math.MaxInt16 {
			res = append(res, int16(v))
		}
	}

	if len(res) > 0 {
		return res
	}

	return fallback
}

// Int32 parse data as int32 or return error on fail
func (c Caster) Int32() (int32, error) {
	if c.IsNil() {
		return 0, c.nilErr()
	}

	if v, ok := asInt64(c.data); ok &&
		v >= math.MinInt32 &&
		v <= math.MaxInt32 {
		return int32(v), nil
	}

	return 0, c.err("failed cast %v as int32!", c.data)
}

// Int32Safe parse data as int32 or return fallback
func (c Caster) Int32Safe(fallback int32) int32 {
	if v, err := c.Int32(); err == nil {
		return v
	}
	return fallback
}

// Int32Slice parse data as []int32 or return fallback
func (c Caster) Int32Slice(fallback []int32) []int32 {
	res := make([]int32, 0)
	for _, item := range c.Slice() {
		if v, ok := asInt64(item); ok &&
			v >= math.MinInt32 &&
			v <= math.MaxInt32 {
			res = append(res, int32(v))
		}
	}

	if len(res) > 0 {
		return res
	}

	return fallback
}

// Int64 parse data as int64 or return error on fail
func (c Caster) Int64() (int64, error) {
	if c.IsNil() {
		return 0, c.nilErr()
	}

	if v, ok := asInt64(c.data); ok {
		return v, nil
	}

	return 0, c.err("failed cast %v as int64!", c.data)
}

// Int64Safe parse data as int64 or return fallback
func (c Caster) Int64Safe(fallback int64) int64 {
	if v, err := c.Int64(); err == nil {
		return v
	}
	return fallback
}

// Int64Slice parse data as []int64 or return fallback
func (c Caster) Int64Slice(fallback []int64) []int64 {
	res := make([]int64, 0)
	for _, item := range c.Slice() {
		if v, ok := asInt64(item); ok {
			res = append(res, v)
		}
	}

	if len(res) > 0 {
		return res
	}

	return fallback
}

// UInt parse data as uint or return error on fail
func (c Caster) UInt() (uint, error) {
	if c.IsNil() {
		return 0, c.nilErr()
	}

	if v, ok := asUint64(c.data); ok && v <= math.MaxUint {
		return uint(v), nil
	}

	return 0, c.err("failed cast %v as uint!", c.data)
}

// UIntSafe parse data as uint or return fallback
func (c Caster) UIntSafe(fallback uint) uint {
	if v, err := c.UInt(); err == nil {
		return v
	}
	return fallback
}

// UIntSlice parse data as []uint or return fallback
func (c Caster) UIntSlice(fallback []uint) []uint {
	res := make([]uint, 0)
	for _, item := range c.Slice() {
		if v, ok := asUint64(item); ok && v <= math.MaxUint {
			res = append(res, uint(v))
		}
	}

	if len(res) > 0 {
		return res
	}

	return fallback
}

// UInt8 parse data as uint8 or return error on fail
func (c Caster) UInt8() (uint8, error) {
	if c.IsNil() {
		return 0, c.nilErr()
	}

	if v, ok := asUint64(c.data); ok && v <= math.MaxUint8 {
		return uint8(v), nil
	}

	return 0, c.err("failed cast %v as uint8!", c.data)
}

// UInt8Safe parse data as uint8 or return fallback
func (c Caster) UInt8Safe(fallback uint8) uint8 {
	if v, err := c.UInt8(); err == nil {
		return v
	}
	return fallback
}

// UInt8Slice parse data as []uint8 or return fallback
func (c Caster) UInt8Slice(fallback []uint8) []uint8 {
	res := make([]uint8, 0)
	for _, item := range c.Slice() {
		if v, ok := asUint64(item); ok && v <= math.MaxUint8 {
			res = append(res, uint8(v))
		}
	}

	if len(res) > 0 {
		return res
	}

	return fallback
}

// UInt16 parse data as uint16 or return error on fail
func (c Caster) UInt16() (uint16, error) {
	if c.IsNil() {
		return 0, c.nilErr()
	}

	if v, ok := asUint64(c.data); ok && v <= math.MaxUint16 {
		return uint16(v), nil
	}

	return 0, c.err("failed cast %v as uint16!", c.data)
}

// UInt16Safe parse data as uint16 or return fallback
func (c Caster) UInt16Safe(fallback uint16) uint16 {
	if v, err := c.UInt16(); err == nil {
		return v
	}
	return fallback
}

// UInt16Slice parse data as []uint16 or return fallback
func (c Caster) UInt16Slice(fallback []uint16) []uint16 {
	res := make([]uint16, 0)
	for _, item := range c.Slice() {
		if v, ok := asUint64(item); ok && v <= math.MaxUint16 {
			res = append(res, uint16(v))
		}
	}

	if len(res) > 0 {
		return res
	}

	return fallback
}

// UInt32 parse data as uint32 or return error on fail
func (c Caster) UInt32() (uint32, error) {
	if c.IsNil() {
		return 0, c.nilErr()
	}

	if v, ok := asUint64(c.data); ok && v <= math.MaxUint32 {
		return uint32(v), nil
	}

	return 0, c.err("failed cast %v as uint32!", c.data)
}

// UInt32Safe parse data as uint32 or return fallback
func (c Caster) UInt32Safe(fallback uint32) uint32 {
	if v, err := c.UInt32(); err == nil {
		return v
	}
	return fallback
}

// UInt32Slice parse data as []uint32 or return fallback
func (c Caster) UInt32Slice(fallback []uint32) []uint32 {
	res := make([]uint32, 0)
	for _, item := range c.Slice() {
		if v, ok := asUint64(item); ok && v <= math.MaxUint32 {
			res = append(res, uint32(v))
		}
	}

	if len(res) > 0 {
		return res
	}

	return fallback
}

// UInt64 parse data as uint64 or return error on fail
func (c Caster) UInt64() (uint64, error) {
	if c.IsNil() {
		return 0, c.nilErr()
	}

	if v, ok := asUint64(c.data); ok {
		return v, nil
	}

	return 0, c.err("failed cast %v as uint64!", c.data)
}

// UInt64Safe parse data as uint64 or return fallback
func (c Caster) UInt64Safe(fallback uint64) uint64 {
	if v, err := c.UInt64(); err == nil {
		return v
	}
	return fallback
}

// UInt64Slice parse data as []uint64 or return fallback
func (c Caster) UInt64Slice(fallback []uint64) []uint64 {
	res := make([]uint64, 0)
	for _, item := range c.Slice() {
		if v, ok := asUint64(item); ok {
			res = append(res, v)
		}
	}

	if len(res) > 0 {
		return res
	}

	return fallback
}

// Float32 parse data as float32 or return error on fail
func (c Caster) Float32() (float32, error) {
	if c.IsNil() {
		return 0, c.nilErr()
	}

	if v, ok := asFloat32(c.data); ok {
		return v, nil
	}

	return 0, c.err("failed cast %v as float32!", c.data)
}

// Float32Safe parse data as float32 or return fallback
func (c Caster) Float32Safe(fallback float32) float32 {
	if v, err := c.Float32(); err == nil {
		return v
	}
	return fallback
}

// Float32Slice parse data as []float32 or return fallback
func (c Caster) Float32Slice(fallback []float32) []float32 {
	res := make([]float32, 0)
	for _, item := range c.Slice() {
		if v, ok := asFloat32(item); ok {
			res = append(res, v)
		}
	}

	if len(res) > 0 {
		return res
	}

	return fallback
}

// Float64 parse data as float64 or return error on fail
func (c Caster) Float64() (float64, error) {
	if c.IsNil() {
		return 0, c.nilErr()
	}

	if v, ok := asFloat64(c.data); ok {
		return v, nil
	}

	return 0, c.err("failed cast %v as float64!", c.data)
}

// Float64Safe parse data as float64 or return fallback
func (c Caster) Float64Safe(fallback float64) float64 {
	if v, err := c.Float64(); err == nil {
		return v
	}
	return fallback
}

// Float64Slice parse data as []float64 or return fallback
func (c Caster) Float64Slice(fallback []float64) []float64 {
	res := make([]float64, 0)
	for _, item := range c.Slice() {
		if v, ok := asFloat64(item); ok {
			res = append(res, v)
		}
	}

	if len(res) > 0 {
		return res
	}

	return fallback
}

// String parse data as string or return error on fail
func (c Caster) String() (string, error) {
	if c.IsNil() {
		return "", c.nilErr()
	}

	if v, ok := asString(c.data); ok {
		return v, nil
	}

	return "", c.err("failed cast %v as string!", c.data)
}

// StringSafe parse data as string or return fallback
func (c Caster) StringSafe(fallback string) string {
	if v, err := c.String(); err == nil {
		return v
	}
	return fallback
}

// StringSlice parse data as []string or return fallback
func (c Caster) StringSlice(fallback []string) []string {
	res := make([]string, 0)
	for _, item := range c.Slice() {
		if v, ok := asString(item); ok {
			res = append(res, v)
		}
	}

	if len(res) > 0 {
		return res
	}

	return fallback
}
