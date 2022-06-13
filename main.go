package caster

// NewCaster create new caster
func NewCaster(data any) Caster {
	return Caster{
		data: data,
	}
}

// IsNilErr check if caster nil error
func IsNilErr(err error) bool {
	return isErrorOf("Caster.Nil", err)
}

// Assert try to assert value to type T
func Assert[T any](v any) (T, bool) {
	_v, ok := v.(T)
	return _v, ok
}

// AssertSafe try to assert value to type T or return fallback
func AssertSafe[T any](v any, fallback T) T {
	if _v, ok := v.(T); ok {
		return _v
	}
	return fallback
}
