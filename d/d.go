// Package d deferences pointers to their values in a nil-safe way.
// If a pointer references nil, and the non-pointer version of its type
// does not allow nil values (error, for instance, may be nil) then the
// zero value of that type is returned.
//
// The letter 'd' stands for deference
package d

import "time"

// Predeclared Identifiers

func Bool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

func Byte(b *byte) byte {
	if b == nil {
		return 0
	}
	return *b
}

func Complex64(c *complex64) complex64 {
	if c == nil {
		return 0
	}
	return *c
}

func Complex128(c *complex128) complex128 {
	if c == nil {
		return 0
	}
	return *c
}

func Error(c *error) error {
	if c == nil {
		return nil
	}
	return *c
}

func Float32(f *float32) float32 {
	if f == nil {
		return 0
	}
	return *f
}

func Float64(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

func Int(n *int) int {
	if n == nil {
		return 0
	}
	return *n
}

func Int8(n *int8) int8 {
	if n == nil {
		return 0
	}
	return *n
}

func Int16(n *int16) int16 {
	if n == nil {
		return 0
	}
	return *n
}

func Int32(n *int32) int32 {
	if n == nil {
		return 0
	}
	return *n
}

func Int64(n *int64) int64 {
	if n == nil {
		return 0
	}
	return *n
}

func Rune(r *rune) rune {
	if r == nil {
		return 0
	}
	return *r
}

func String(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func Uint(n *uint) uint {
	if n == nil {
		return 0
	}
	return *n
}

func Uint8(n *uint8) uint8 {
	if n == nil {
		return 0
	}
	return *n
}

func Uint16(n *uint16) uint16 {
	if n == nil {
		return 0
	}
	return *n
}

func Uint32(n *uint32) uint32 {
	if n == nil {
		return 0
	}
	return *n
}

func Uint64(n *uint64) uint64 {
	if n == nil {
		return 0
	}
	return *n
}

func Uintptr(n *uintptr) uintptr {
	if n == nil {
		return 0
	}
	return *n
}

// Standard types

func Time(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}
