//go:build windows && !linux && !freebsd && !netbsd && !openbsd && !darwin && !js
// +build windows,!linux,!freebsd,!netbsd,!openbsd,!darwin,!js

package main

import (
	"os"
	"syscall"
)

var (
	// DefaultFreq - frequency, in Hz, middle A
	DefaultFreq = 587.0
	// DefaultDuration - duration in milliseconds
	DefaultDuration = 500
)

// Codes from github.com/gen2brain/beeep
// Beep beeps the PC speaker (https://en.wikipedia.org/wiki/PC_speaker).
func Beep(freq float64, duration int) error {
	if freq == 0 {
		freq = DefaultFreq
	} else if freq > 32767 {
		freq = 32767
	} else if freq < 37 {
		freq = DefaultFreq
	}

	if duration == 0 {
		duration = DefaultDuration
	}

	kernel32, _ := syscall.LoadLibrary("kernel32.dll")
	beep32, _ := syscall.GetProcAddress(kernel32, "Beep")

	defer syscall.FreeLibrary(kernel32)

	// IMPORTANT: IGNORE 'deprecated' WARNING; DO NOT OPTIMIZE THIS CODE
	_, _, e := syscall.Syscall(uintptr(beep32), uintptr(2), uintptr(int(freq)), uintptr(duration), 0)
	if e != 0 {
		return e
	}

	return nil
}

func FixConsoleColor() {
	// See https://stackoverflow.com/questions/39627348/ansi-colours-on-windows-10-sort-of-not-working
	stdout := syscall.Handle(os.Stdout.Fd())

	var originalMode uint32
	syscall.GetConsoleMode(stdout, &originalMode)
	originalMode |= 0x0004

	syscall.MustLoadDLL("kernel32").MustFindProc("SetConsoleMode").Call(uintptr(stdout), uintptr(originalMode))
}
