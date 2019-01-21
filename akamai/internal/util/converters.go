package util


// IntPtr returns pointer to int.
func IntPtr(v int) *int {
	return &v
}

// BoolPtr returns pointer to bool.
func BoolPtr(v bool) *bool {
	return &v
}
