package caster_test

import (
	"testing"

	"github.com/bopher/caster"
)

func TestIsNil(t *testing.T) {
	c := caster.NewCaster(nil)
	if !c.IsNil() {
		t.Fatal("failed")
	}
}

func TestSlice(t *testing.T) {
	c := caster.NewCaster([]interface{}{false, "A", 1.234})
	if len(c.Slice()) != 3 {
		t.Fatal("failed")
	}
}

func TestBool(t *testing.T) {
	c := caster.NewCaster(true)
	if c.BoolSafe(false) == false {
		t.Fatal("failed")
	}
}

func TestBoolSlice(t *testing.T) {
	c := caster.NewCaster([]interface{}{false, true, true})
	if c.BoolSlice(nil) == nil {
		t.Fatal("failed")
	}
}

func TestInt(t *testing.T) {
	c := caster.NewCaster(5331)
	if c.IntSafe(1) == 1 {
		t.Fatal("failed")
	}
}

func TestIntSlice(t *testing.T) {
	c := caster.NewCaster([]interface{}{4, 5, 6})
	if c.IntSlice(nil) == nil {
		t.Fatal("failed")
	}
}

func TestIn8(t *testing.T) {
	c := caster.NewCaster(300)
	if c.Int8Safe(2) != 2 {
		t.Fatal("failed")
	}
}

func TestInt8Slice(t *testing.T) {
	c := caster.NewCaster([]interface{}{25, 10, 12})
	if c.Int8Slice(nil) == nil {
		t.Fatal("failed")
	}
}

func TestInt16(t *testing.T) {
	c := caster.NewCaster(120)
	if c.Int16Safe(3) == 3 {
		t.Fatal("failed")
	}
}

func TestInt16Slice(t *testing.T) {
	c := caster.NewCaster([]interface{}{10, 20, 30})
	if c.Int16Slice(nil) == nil {
		t.Fatal("failed")
	}
}

func TestInt32(t *testing.T) {
	c := caster.NewCaster(555)
	if c.Int32Safe(9) == 9 {
		t.Fatal("failed")
	}
}

func TestInt32Slice(t *testing.T) {
	c := caster.NewCaster([]interface{}{5, 9, 100})
	if c.Int32Slice(nil) == nil {
		t.Fatal("failed")
	}
}

func TestInt64(t *testing.T) {
	c := caster.NewCaster(120.99)
	if c.Int64Safe(8) != 8 {
		t.Fatal("failed")
	}
}

func TestInt64Slice(t *testing.T) {
	c := caster.NewCaster([]interface{}{2131, 55555})
	if c.Int64Slice(nil) == nil {
		t.Fatal("failed")
	}
}

func TestUInt(t *testing.T) {
	c := caster.NewCaster(-100)
	if c.UIntSafe(4) != 4 {
		t.Fatal("failed")
	}
}

func TestUIntSlice(t *testing.T) {
	c := caster.NewCaster([]interface{}{100, -20, 30})
	if len(c.UIntSlice(nil)) != 2 {
		t.Fatal("failed")
	}
}

func TestUInt8(t *testing.T) {
	c := caster.NewCaster(255)
	if c.UInt8Safe(3) == 3 {
		t.Fatal("failed")
	}
}

func TestUInt8Slice(t *testing.T) {
	c := caster.NewCaster([]interface{}{100, -20, 300})
	if len(c.UInt8Slice(nil)) != 1 {
		t.Fatal("failed")
	}
}

func TestUInt16(t *testing.T) {
	c := caster.NewCaster(255)
	if c.UInt16Safe(3) == 3 {
		t.Fatal("failed")
	}
}

func TestUInt16Slice(t *testing.T) {
	c := caster.NewCaster([]interface{}{100, -20, 300})
	if len(c.UInt16Slice(nil)) != 2 {
		t.Fatal("failed")
	}
}

func TestUInt32(t *testing.T) {
	c := caster.NewCaster(255)
	if c.UInt32Safe(3) == 3 {
		t.Fatal("failed")
	}
}

func TestUInt32Slice(t *testing.T) {
	c := caster.NewCaster([]interface{}{100, -20, 300})
	if len(c.UInt32Slice(nil)) != 2 {
		t.Fatal("failed")
	}
}

func TestUInt64(t *testing.T) {
	c := caster.NewCaster(255)
	if c.UInt64Safe(3) == 3 {
		t.Fatal("failed")
	}
}

func TestUInt64Slice(t *testing.T) {
	c := caster.NewCaster([]interface{}{-1, -20, 300})
	if len(c.UInt64Slice(nil)) != 1 {
		t.Fatal("failed")
	}
}

func TestFloat32(t *testing.T) {
	c := caster.NewCaster("1234.12aaa")
	if c.Float32Safe(11) != 11 {
		t.Fatal("failed")
	}
}

func TestFloat32Slice(t *testing.T) {
	c := caster.NewCaster([]interface{}{2.3, 700, -19.22})
	if len(c.Float32Slice(nil)) != 3 {
		t.Fatal("failed")
	}
}

func TestFloat64(t *testing.T) {
	c := caster.NewCaster(33.444123123)
	if c.Float64Safe(4) == 4 {
		t.Fatal("failed")
	}
}

func TestFloat64Slice(t *testing.T) {
	c := caster.NewCaster([]interface{}{10, true, 20.11})
	if len(c.Float64Slice(nil)) != 2 {
		t.Fatal("failed")
	}
}

func TestString(t *testing.T) {
	c := caster.NewCaster(123)
	if c.StringSafe("hi") != "hi" {
		t.Fatal("failed")
	}
}

func TestStringSlice(t *testing.T) {
	c := caster.NewCaster([]interface{}{"John", "Jack", "Jimmy"})
	if len(c.StringSlice(nil)) != 3 {
		t.Fatal("failed")
	}
}
