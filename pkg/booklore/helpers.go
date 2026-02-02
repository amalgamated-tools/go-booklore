package booklore

import "time"

func StrToPtr(s string) *string {
	return &s
}

// Nil-safe accessor functions for pointer types
// These return zero values instead of requiring nil checks

// GetString returns the dereferenced string or empty string if nil
func GetString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// GetInt64 returns the dereferenced int64 or 0 if nil
func GetInt64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

// GetInt32 returns the dereferenced int32 or 0 if nil
func GetInt32(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}

// GetBool returns the dereferenced bool or false if nil
func GetBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

// GetFloat64 returns the dereferenced float64 or 0 if nil
func GetFloat64(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

// GetTime returns the dereferenced time.Time or zero time if nil
func GetTime(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}
