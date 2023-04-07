//go:build !plan9 && !windows && !wasm && !js

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */

package conn

import (
	"golang.org/x/sys/unix"
	"syscall"
)

func init() {

	controlFns = append(controlFns,
		func(network, address string, c syscall.RawConn) error {
			return c.Control(func(fd uintptr) {
				_ = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEADDR, 1)
			})
		},
	)
	controlFns = append(controlFns,
		func(network, address string, c syscall.RawConn) error {
			return c.Control(func(fd uintptr) {
				_ = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
			})
		},
	)

}
