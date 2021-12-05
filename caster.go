package caster

import (
	"math"
)

type Caster struct {
	data interface{}
}

func (Caster) err(pattern string, params ...interface{}) error {
	return taggedError([]string{"Caster"}, pattern, params...)
}

func (Caster) nilErr() error {
	return taggedError([]string{"Caster", "Caster.Nil"}, "value is nil!")
}

// IsNil check if value of caster is nil
func (this Caster) IsNil() bool {
	return this.data == nil
}

// Interface get raw value
func (this Caster) Interface() interface{} {
	return this.data
}

// Slice parse data as []interface{}
func (this Caster) Slice() []interface{} {
	return asSlice(this.data)
}

// Bool parse data as boolean or return error on fail
func (this Caster) Bool() (bool, error) {
	if this.IsNil() {
		return false, this.nilErr()
	}

	if v, ok := asBool(this.data); ok {
		return v, nil
	}
	return false, this.err(`failed cast "%v" as bool!`, this.data)
}

// BoolSafe parse data as boolean or return fallback
func (this Caster) BoolSafe(fallback bool) bool {
	if v, err := this.Bool(); err == nil {
		return v
	}
	return fallback
}

// BoolSlice parse data as []bool or return fallback
func (this Caster) BoolSlice(fallback []bool) []bool {
	res := make([]bool, 0)
	for _, item := range this.Slice() {
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
func (this Caster) Int() (int, error) {
	if this.IsNil() {
		return 0, this.nilErr()
	}

	if v, ok := asInt64(this.data); ok &&
		v >= math.MinInt &&
		v <= math.MaxInt {
		return int(v), nil
	}

	return 0, this.err("failed cast %v as int!", this.data)
}

// IntSafe parse data as int or return fallback
func (this Caster) IntSafe(fallback int) int {
	if v, err := this.Int(); err == nil {
		return v
	}
	return fallback
}

// IntSlice parse data as []int or return fallback
func (this Caster) IntSlice(fallback []int) []int {
	res := make([]int, 0)
	for _, item := range this.Slice() {
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
func (this Caster) Int8() (int8, error) {
	if this.IsNil() {
		return 0, this.nilErr()
	}

	if v, ok := asInt64(this.data); ok &&
		v >= math.MinInt8 &&
		v <= math.MaxInt8 {
		return int8(v), nil
	}

	return 0, this.err("failed cast %v as int8!", this.data)
}

// Int8Safe parse data as int8 or return fallback
func (this Caster) Int8Safe(fallback int8) int8 {
	if v, err := this.Int8(); err == nil {
		return v
	}
	return fallback
}

// Int8Slice parse data as []int8 or return fallback
func (this Caster) Int8Slice(fallback []int8) []int8 {
	res := make([]int8, 0)
	for _, item := range this.Slice() {
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
func (this Caster) Int16() (int16, error) {
	if this.IsNil() {
		return 0, this.nilErr()
	}

	if v, ok := asInt64(this.data); ok &&
		v >= math.MinInt16 &&
		v <= math.MaxInt16 {
		return int16(v), nil
	}

	return 0, this.err("failed cast %v as int16!", this.data)
}

// Int16Safe parse data as int16  or return fallback
func (this Caster) Int16Safe(fallback int16) int16 {
	if v, err := this.Int16(); err == nil {
		return v
	}
	return fallback
}

// Int16Slice parse data as []int16 or return fallback
func (this Caster) Int16Slice(fallback []int16) []int16 {
	res := make([]int16, 0)
	for _, item := range this.Slice() {
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
func (this Caster) Int32() (int32, error) {
	if this.IsNil() {
		return 0, this.nilErr()
	}

	if v, ok := asInt64(this.data); ok &&
		v >= math.MinInt32 &&
		v <= math.MaxInt32 {
		return int32(v), nil
	}

	return 0, this.err("failed cast %v as int32!", this.data)
}

// Int32Safe parse data as int32 or return fallback
func (this Caster) Int32Safe(fallback int32) int32 {
	if v, err := this.Int32(); err == nil {
		return v
	}
	return fallback
}

// Int32Slice parse data as []int32 or return fallback
func (this Caster) Int32Slice(fallback []int32) []int32 {
	res := make([]int32, 0)
	for _, item := range this.Slice() {
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
func (this Caster) Int64() (int64, error) {
	if this.IsNil() {
		return 0, this.nilErr()
	}

	if v, ok := asInt64(this.data); ok {
		return v, nil
	}

	return 0, this.err("failed cast %v as int64!", this.data)
}

// Int64Safe parse data as int64 or return fallback
func (this Caster) Int64Safe(fallback int64) int64 {
	if v, err := this.Int64(); err == nil {
		return v
	}
	return fallback
}

// Int64Slice parse data as []int64 or return fallback
func (this Caster) Int64Slice(fallback []int64) []int64 {
	res := make([]int64, 0)
	for _, item := range this.Slice() {
		if v, ok := asInt64(item); ok &&
			v >= math.MinInt64 &&
			v <= math.MaxInt64 {
			res = append(res, v)
		}
	}

	if len(res) > 0 {
		return res
	}

	return fallback
}

// UInt parse data as uint or return error on fail
func (this Caster) UInt() (uint, error) {
	if this.IsNil() {
		return 0, this.nilErr()
	}

	if v, ok := asUint64(this.data); ok && v <= math.MaxUint {
		return uint(v), nil
	}

	return 0, this.err("failed cast %v as uint!", this.data)
}

// UIntSafe parse data as uint or return fallback
func (this Caster) UIntSafe(fallback uint) uint {
	if v, err := this.UInt(); err == nil {
		return v
	}
	return fallback
}

// UIntSlice parse data as []uint or return fallback
func (this Caster) UIntSlice(fallback []uint) []uint {
	res := make([]uint, 0)
	for _, item := range this.Slice() {
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
func (this Caster) UInt8() (uint8, error) {
	if this.IsNil() {
		return 0, this.nilErr()
	}

	if v, ok := asUint64(this.data); ok && v <= math.MaxUint8 {
		return uint8(v), nil
	}

	return 0, this.err("failed cast %v as uint8!", this.data)
}

// UInt8Safe parse data as uint8 or return fallback
func (this Caster) UInt8Safe(fallback uint8) uint8 {
	if v, err := this.UInt8(); err == nil {
		return v
	}
	return fallback
}

// UInt8Slice parse data as []uint8 or return fallback
func (this Caster) UInt8Slice(fallback []uint8) []uint8 {
	res := make([]uint8, 0)
	for _, item := range this.Slice() {
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
func (this Caster) UInt16() (uint16, error) {
	if this.IsNil() {
		return 0, this.nilErr()
	}

	if v, ok := asUint64(this.data); ok && v <= math.MaxUint16 {
		return uint16(v), nil
	}

	return 0, this.err("failed cast %v as uint16!", this.data)
}

// UInt16Safe parse data as uint16 or return fallback
func (this Caster) UInt16Safe(fallback uint16) uint16 {
	if v, err := this.UInt16(); err == nil {
		return v
	}
	return fallback
}

// UInt16Slice parse data as []uint16 or return fallback
func (this Caster) UInt16Slice(fallback []uint16) []uint16 {
	res := make([]uint16, 0)
	for _, item := range this.Slice() {
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
func (this Caster) UInt32() (uint32, error) {
	if this.IsNil() {
		return 0, this.nilErr()
	}

	if v, ok := asUint64(this.data); ok && v <= math.MaxUint32 {
		return uint32(v), nil
	}

	return 0, this.err("failed cast %v as uint32!", this.data)
}

// UInt32Safe parse data as uint32 or return fallback
func (this Caster) UInt32Safe(fallback uint32) uint32 {
	if v, err := this.UInt32(); err == nil {
		return v
	}
	return fallback
}

// UInt32Slice parse data as []uint32 or return fallback
func (this Caster) UInt32Slice(fallback []uint32) []uint32 {
	res := make([]uint32, 0)
	for _, item := range this.Slice() {
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
func (this Caster) UInt64() (uint64, error) {
	if this.IsNil() {
		return 0, this.nilErr()
	}

	if v, ok := asUint64(this.data); ok {
		return v, nil
	}

	return 0, this.err("failed cast %v as uint64!", this.data)
}

// UInt64Safe parse data as uint64 or return fallback
func (this Caster) UInt64Safe(fallback uint64) uint64 {
	if v, err := this.UInt64(); err == nil {
		return v
	}
	return fallback
}

// UInt64Slice parse data as []uint64 or return fallback
func (this Caster) UInt64Slice(fallback []uint64) []uint64 {
	res := make([]uint64, 0)
	for _, item := range this.Slice() {
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
func (this Caster) Float32() (float32, error) {
	if this.IsNil() {
		return 0, this.nilErr()
	}

	if v, ok := asFloat32(this.data); ok {
		return v, nil
	}

	return 0, this.err("failed cast %v as float32!", this.data)
}

// Float32Safe parse data as float32 or return fallback
func (this Caster) Float32Safe(fallback float32) float32 {
	if v, err := this.Float32(); err == nil {
		return v
	}
	return fallback
}

// Float32Slice parse data as []float32 or return fallback
func (this Caster) Float32Slice(fallback []float32) []float32 {
	res := make([]float32, 0)
	for _, item := range this.Slice() {
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
func (this Caster) Float64() (float64, error) {
	if this.IsNil() {
		return 0, this.nilErr()
	}

	if v, ok := asFloat64(this.data); ok {
		return v, nil
	}

	return 0, this.err("failed cast %v as float64!", this.data)
}

// Float64Safe parse data as float64 or return fallback
func (this Caster) Float64Safe(fallback float64) float64 {
	if v, err := this.Float64(); err == nil {
		return v
	}
	return fallback
}

// Float64Slice parse data as []float64 or return fallback
func (this Caster) Float64Slice(fallback []float64) []float64 {
	res := make([]float64, 0)
	for _, item := range this.Slice() {
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
func (this Caster) String() (string, error) {
	if this.IsNil() {
		return "", this.nilErr()
	}

	if v, ok := asString(this.data); ok {
		return v, nil
	}

	return "", this.err("failed cast %v as string!", this.data)
}

// StringSafe parse data as string or return fallback
func (this Caster) StringSafe(fallback string) string {
	if v, err := this.String(); err == nil {
		return v
	}
	return fallback
}

// StringSlice parse data as []string or return fallback
func (this Caster) StringSlice(fallback []string) []string {
	res := make([]string, 0)
	for _, item := range this.Slice() {
		if v, ok := asString(item); ok {
			res = append(res, v)
		}
	}

	if len(res) > 0 {
		return res
	}

	return fallback
}
