//go:build !windows
// +build !windows

package transport

import (
	"fmt"
	"syscall"

	"golang.org/x/sys/unix"
)

func setReusePort(network, address string, conn syscall.RawConn) error {
	fmt.Println("DANISH HERE")
	return conn.Control(func(fd uintptr) {
		syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, unix.SO_REUSEPORT, 1)
	})
}

func setReuseAddress(network, address string, conn syscall.RawConn) error {
	return conn.Control(func(fd uintptr) {
		syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, unix.SO_REUSEADDR, 1)
	})
}
