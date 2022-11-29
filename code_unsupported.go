//go:build !linux && !freebsd && !netbsd && !openbsd && !windows && !darwin
// +build !linux,!freebsd,!netbsd,!openbsd,!windows,!darwin

package main

var (
	// DefaultFreq - frequency, in Hz, middle A
	DefaultFreq = 0.0
	// DefaultDuration - duration in milliseconds
	DefaultDuration = 0
)

// Codes from github.com/gen2brain/beeep
// Beep beeps the PC speaker (https://en.wikipedia.org/wiki/PC_speaker).
func Beep(freq float64, duration int) error {
	return ErrUnsupported
}

func FixConsoleColor() {
	// Empty
}
