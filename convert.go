package safeconv

import "fmt"

type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func ConvertOK[From, To number](v From) (To, bool) {
	vTo := To(v)
	if From(vTo) != v || (v < 0) != (vTo < 0) {
		return 0, false
	}
	return vTo, true
}

func Convert[From, To number](v From) To {
	c, ok := ConvertOK[From, To](v)
	if !ok {
		panic(fmt.Sprintf("%v cannot be represented exactly as %T", v, To(0)))
	}
	return c
}

func Float64OK[T number](v T) (float64, bool) {
	return ConvertOK[T, float64](v)
}

func Float32OK[T number](v T) (float32, bool) {
	return ConvertOK[T, float32](v)
}

func IntOK[T number](v T) (int, bool) {
	return ConvertOK[T, int](v)
}

func Int8OK[T number](v T) (int8, bool) {
	return ConvertOK[T, int8](v)
}

func Int16OK[T number](v T) (int16, bool) {
	return ConvertOK[T, int16](v)
}

func Int32OK[T number](v T) (int32, bool) {
	return ConvertOK[T, int32](v)
}

func Int64OK[T number](v T) (int64, bool) {
	return ConvertOK[T, int64](v)
}

func UintOK[T number](v T) (uint, bool) {
	return ConvertOK[T, uint](v)
}

func Uint8OK[T number](v T) (uint8, bool) {
	return ConvertOK[T, uint8](v)
}

func Uint16OK[T number](v T) (uint16, bool) {
	return ConvertOK[T, uint16](v)
}

func Uint32OK[T number](v T) (uint32, bool) {
	return ConvertOK[T, uint32](v)
}

func Uint64OK[T number](v T) (uint64, bool) {
	return ConvertOK[T, uint64](v)
}

func UintptrOK[T number](v T) (uintptr, bool) {
	return ConvertOK[T, uintptr](v)
}

func Float64[T number](v T) float64 {
	return Convert[T, float64](v)
}

func Float32[T number](v T) float32 {
	return Convert[T, float32](v)
}

func Int[T number](v T) int {
	return Convert[T, int](v)
}

func Int8[T number](v T) int8 {
	return Convert[T, int8](v)
}

func Int16[T number](v T) int16 {
	return Convert[T, int16](v)
}

func Int32[T number](v T) int32 {
	return Convert[T, int32](v)
}

func Int64[T number](v T) int64 {
	return Convert[T, int64](v)
}

func Uint[T number](v T) uint {
	return Convert[T, uint](v)
}

func Uint8[T number](v T) uint8 {
	return Convert[T, uint8](v)
}

func Uint16[T number](v T) uint16 {
	return Convert[T, uint16](v)
}

func Uint32[T number](v T) uint32 {
	return Convert[T, uint32](v)
}

func Uint64[T number](v T) uint64 {
	return Convert[T, uint64](v)
}

func Uintptr[T number](v T) uintptr {
	return Convert[T, uintptr](v)
}
