package caster

// NewCaster create new caster
func NewCaster(data interface{}) Caster {
	return Caster{
		data: data,
	}
}

// IsNilErr check if caster nil error
func IsNilErr(err error) bool {
	return isErrorOf("Caster.Nil", err)
}
