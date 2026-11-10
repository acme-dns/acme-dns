//go:build windows

package main

func setUmask() {
	// umask is not supported on Windows
}
