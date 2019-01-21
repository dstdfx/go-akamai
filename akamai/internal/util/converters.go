package util


// IntPtr returns pointer to int.
func IntPtr(v int) *int {
	return &v
}

// BoolPtr returns pointer to bool.
func BoolPtr(v bool) *bool {
	return &v
}

// StringPtr returns pointer to string.
func StringPtr(v string) *string {
	return &v
}
