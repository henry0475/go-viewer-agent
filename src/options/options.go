package options

// Bool returns a pointer to the string value passed in.
func Bool(b bool) *bool {
	return &b
}

// Int64 returns a pointer to the int64 value passed in.
func Int64(i int64) *int64 {
	return &i
}
