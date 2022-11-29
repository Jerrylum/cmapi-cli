//go:build darwin && !linux && !freebsd && !netbsd && !openbsd && !windows && !js
// +build darwin,!linux,!freebsd,!netbsd,!openbsd,!windows,!js

package main

import (
	"os"
	"os/exec"
)

var (
	// DefaultFreq - frequency, in Hz, middle A
	DefaultFreq = 0.0
	// DefaultDuration - duration in milliseconds
	DefaultDuration = 0
)

// Codes from github.com/gen2brain/beeep
// Beep beeps the PC speaker (https://en.wikipedia.org/wiki/PC_speaker).
func Beep(freq float64, duration int) error {
	osa, err := exec.LookPath("osascript")
	if err != nil {
		// Output the only beep we can
		_, err = os.Stdout.Write([]byte{7})
		return err
	}

	cmd := exec.Command(osa, "-e", `beep`)
	return cmd.Run()
}

func FixConsoleColor() {
	// Empty
}
