# Caster

Error-safe value caster.

## NewCaster

Create new caster instance from value.

```go
// Signature
NewCaster(data interface{}) Caster

// Example
import "github.com/bopher/caster"
c := caster.NewCaster(1234)
```

## IsNilErr

Check if returns error is nil error.

```go
// Signature
IsNilErr(err error) bool

// Example
import "github.com/bopher/caster"
c := caster.NewCaster(nil)
v, err := c.Int()
if caster.IsNilErr(err) {
    // ...
}
```

## Methods

Caster has following methods.

### IsNil

Check if value of caster is nil.

```go
IsNil() bool
```

### Raw

Get raw value.

```go
Raw() interface{}
```

### Bool

Parse data as boolean or return error on fail.

```go
Bool() (bool, error)
```

### BoolSafe

Parse data as boolean or return fallback.

```go
BoolSafe(fallback bool) bool
```

### Int

Parse data as int or return error on fail.

```go
Int() (int, error)
```

### IntSafe

Parse data as int or return fallback.

```go
IntSafe(fallback int) int
```

### Int8

Parse data as int8 or return error on fail.

```go
Int8() (int8, error)
```

### Int8Safe

Parse data as int8 or return fallback.

```go
Int8Safe(fallback int8) int8
```

### Int16

Parse data as int16 or return error on fail.

```go
Int16() (int16, error)
```

### Int16Safe

Parse data as int16 or return fallback.

```go
Int16Safe(fallback int16) int16
```

### Int32

Parse data as int32 or return error on fail.

```go
Int32() (int32, error)
```

### Int32Safe

Parse data as int32 or return fallback.

```go
Int32Safe(fallback int32) int32
```

### Int64

Parse data as int64 or return error on fail.

```go
Int64() (int64, error)
```

### Int64Safe

Parse data as int64 or return fallback.

```go
Int64Safe(fallback int64) int64
```

### UInt

Parse data as uint or return error on fail.

```go
UInt() (uint, error)
```

### UIntSafe

Parse data as uint or return fallback.

```go
UIntSafe(fallback uint) uint
```

### UInt8

Parse data as uint8 or return error on fail.

```go
UInt8() (uint8, error)
```

### UInt8Safe

Parse data as uint8 or return fallback.

```go
UInt8Safe(fallback uint8) uint8
```

### UInt16

Parse data as uint16 or return error on fail.

```go
UInt16() (uint16, error)
```

### UInt16Safe

Parse data as uint16 or return fallback.

```go
UInt16Safe(fallback uint16) uint16
```

### UInt32

Parse data as uint32 or return error on fail.

```go
UInt32() (uint32, error)
```

### UInt32Safe

Parse data as uint32 or return fallback.

```go
UInt32Safe(fallback uint32) uint32
```

### UInt64

Parse data as uint64 or return error on fail.

```go
UInt64() (uint64, error)
```

### UInt64Safe

Parse data as uint64 or return fallback.

```go
UInt64Safe(fallback uint64) uint64
```

### Float32

Parse data as float32 or return error on fail.

```go
Float32() (float32, error)
```

### Float32Safe

Parse data as float32 or return fallback.

```go
Float32Safe(fallback float32) float32
```

### Float64

Parse data as float64 or return error on fail.

```go
Float64() (float64, error)
```

### Float64Safe

Parse data as float64 or return fallback.

```go
Float64Safe(fallback float64) float64
```

### String

Parse data as string or return error on fail.

```go
String() (string, error)
```

### StringSafe

Parse data as string or return fallback.

```go
StringSafe(fallback string) string
```

### Interface

Parse data as `interface{}` or return error on fail.

```go
Interface() (interface{}, error)
```
