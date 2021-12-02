package caster

import (
	"fmt"
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

// Raw get raw value
func (this Caster) Raw() interface{} {
	return this.data
}

// Bool parse data as boolean or return error on fail
func (this Caster) Bool() (bool, error) {
	if this.IsNil() {
		return false, this.nilErr()
	}

	if r, ok := asBool(this.data); ok {
		return r, nil
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

// Int parse data as int or return error on fail
func (this Caster) Int() (int, error) {
	if this.IsNil() {
		return 0, this.nilErr()
	}

	if r, ok := asInt64(this.data); ok &&
		r >= math.MinInt &&
		r <= math.MaxInt {
		return int(r), nil
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

// Int8 parse data as int8 or return error on fail
func (this Caster) Int8() (int8, error) {
	if this.IsNil() {
		return 0, this.nilErr()
	}

	if r, ok := asInt64(this.data); ok &&
		r >= math.MinInt8 &&
		r <= math.MaxInt8 {
		return int8(r), nil
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

// Int16 parse data as int16 or return error on fail
func (this Caster) Int16() (int16, error) {
	if this.IsNil() {
		return 0, this.nilErr()
	}

	if r, ok := asInt64(this.data); ok &&
		r >= math.MinInt16 &&
		r <= math.MaxInt16 {
		return int16(r), nil
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

// Int32 parse data as int32 or return error on fail
func (this Caster) Int32() (int32, error) {
	if this.IsNil() {
		return 0, this.nilErr()
	}

	if r, ok := asInt64(this.data); ok &&
		r >= math.MinInt32 &&
		r <= math.MaxInt32 {
		return int32(r), nil
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

// Int64 parse data as int64 or return error on fail
func (this Caster) Int64() (int64, error) {
	if this.IsNil() {
		return 0, this.nilErr()
	}

	if r, ok := asInt64(this.data); ok &&
		r >= math.MinInt64 &&
		r <= math.MaxInt64 {
		return r, nil
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

// UInt parse data as uint or return error on fail
func (this Caster) UInt() (uint, error) {
	if this.IsNil() {
		return 0, this.nilErr()
	}

	if r, ok := asUint64(this.data); ok &&
		r >= 0 &&
		r <= math.MaxUint {
		return uint(r), nil
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

// UInt8 parse data as uint8 or return error on fail
func (this Caster) UInt8() (uint8, error) {
	if this.IsNil() {
		return 0, this.nilErr()
	}

	if r, ok := asUint64(this.data); ok &&
		r >= 0 &&
		r <= math.MaxUint8 {
		return uint8(r), nil
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

// UInt16 parse data as uint16 or return error on fail
func (this Caster) UInt16() (uint16, error) {
	if this.IsNil() {
		return 0, this.nilErr()
	}

	if r, ok := asUint64(this.data); ok &&
		r >= 0 &&
		r <= math.MaxUint16 {
		return uint16(r), nil
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

// UInt32 parse data as uint32 or return error on fail
func (this Caster) UInt32() (uint32, error) {
	if this.IsNil() {
		return 0, this.nilErr()
	}

	if r, ok := asUint64(this.data); ok &&
		r >= 0 &&
		r <= math.MaxUint32 {
		return uint32(r), nil
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

// UInt64 parse data as uint64 or return error on fail
func (this Caster) UInt64() (uint64, error) {
	if this.IsNil() {
		return 0, this.nilErr()
	}

	if r, ok := asUint64(this.data); ok &&
		r >= 0 &&
		r <= math.MaxUint64 {
		return r, nil
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

// Float32 parse data as float32 or return error on fail
func (this Caster) Float32() (float32, error) {
	if this.IsNil() {
		return 0, this.nilErr()
	}

	if r, ok := asFloat32(this.data); ok {
		return float32(r), nil
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

// Float64 parse data as float64 or return error on fail
func (this Caster) Float64() (float64, error) {
	if this.IsNil() {
		return 0, this.nilErr()
	}

	if r, ok := asFloat64(this.data); ok {
		return r, nil
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

// String parse data as string or return error on fail
func (this Caster) String() (string, error) {
	if this.IsNil() {
		return "", this.nilErr()
	}

	return fmt.Sprint(this.data), nil
}

// StringSafe parse data as string or return fallback
func (this Caster) StringSafe(fallback string) string {
	if v, err := this.String(); err == nil {
		return v
	}
	return fallback
}

// Interface parse data as interface{} or return error on fail
func (this Caster) Interface() (interface{}, error) {
	if this.IsNil() {
		return "", this.nilErr()
	}

	return this.data, nil
}
