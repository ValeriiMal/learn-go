package logs

import "strings"

// Application identifies the application emitting the given log.
func Application(log string) string {
	var app string = "default"
	for _, r := range log {
		if r == '❗' {
			app = "recommendation"
			break
		} else if r == '🔍' {
			app = "search"
			break
		} else if r == '☀' {
			app = "weather"
			break
		}
	}
	return app
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	var length int
	for range log {
		length++
	}
	return length <= limit
}
