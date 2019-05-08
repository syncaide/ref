// Package ref implements primitive type pointer de-referencing function.
package ref

func Bool(value bool) *bool {
	return &value
}

func Byte(value byte) *byte {
	return &value
}

func Complex64(value complex64) *complex64 {
	return &value
}

func Complex128(value complex128) *complex128 {
	return &value
}

func Float32(value float32) *float32 {
	return &value
}

func Float64(value float64) *float64 {
	return &value
}

func Int(value int) *int {
	return &value
}

func Int8(value int8) *int8 {
	return &value
}

func Int16(value int16) *int16 {
	return &value
}

func Int32(value int32) *int32 {
	return &value
}

func Int64(value int64) *int64 {
	return &value
}

func Rune(value rune) *rune {
	return &value
}

func String(value string) *string {
	return &value
}

func Uint(value uint) *uint {
	return &value
}

func Uint8(value uint8) *uint8 {
	return &value
}

func Uint16(value uint16) *uint16 {
	return &value
}

func Uint32(value uint32) *uint32 {
	return &value
}

func Uint64(value uint64) *uint64 {
	return &value
}

func Slice(value []interface{}) *[]interface{} {
	return &value
}
