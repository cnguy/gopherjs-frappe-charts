package utils

// Btoi is a simple helper to translate booleans into their integer counterparts.
// This is made because frappe-charts uses 0's and 1's, but this library
// also provides helper functions that allow you to do things such as:
// chartArgs.SetIsNavigable(true)
// instead of:
// chartArgs.IsNavigable = 1
func Btoi(b bool) int {
	if !b {
		return 0
	}
	return 1
}
