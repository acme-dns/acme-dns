//go:build !windows

package main

import "syscall"

func setUmask() {
	syscall.Umask(0077)
}
