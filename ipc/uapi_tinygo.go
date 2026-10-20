//go:build tinygo && (linux || darwin || freebsd || openbsd)

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2025 WireGuard LLC. All Rights Reserved.
 */

package ipc

import (
	"errors"
	"net"
	"os"

	"golang.org/x/sys/unix"
)

// Same values as uapi_unix.go; TinyGo's net package lacks the unix socket
// APIs (net.ListenUnix, net.FileListener) the UAPI listener is built on.
const (
	IpcErrorIO        = -int64(unix.EIO)
	IpcErrorProtocol  = -int64(unix.EPROTO)
	IpcErrorInvalid   = -int64(unix.EINVAL)
	IpcErrorPortInUse = -int64(unix.EADDRINUSE)
	IpcErrorUnknown   = -55 // ENOANO
)

var errUAPIUnsupported = errors.New("ipc: UAPI socket unsupported under TinyGo")

func UAPIOpen(name string) (*os.File, error) {
	return nil, errUAPIUnsupported
}

func UAPIListen(name string, file *os.File) (net.Listener, error) {
	return nil, errUAPIUnsupported
}
