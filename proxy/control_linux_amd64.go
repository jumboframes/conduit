/*
 * Apache License 2.0
 *
 * Copyright (c) 2022, Austin Zhai
 * All rights reserved.
 */
// +build linux,amd64

package proxy

import (
	"syscall"

	"github.com/jumboframes/conduit/pkg/log"

	"golang.org/x/sys/unix"
)

func control(network, address string, conn syscall.RawConn) error {
	var operr, err error
	err = conn.Control(func(fd uintptr) {
		operr = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
		if operr != nil {
			return
		}
		operr = syscall.SetsockoptInt(int(fd), unix.SOL_SOCKET, syscall.SO_MARK, 5053)
		if operr != nil {
			return
		}
	})
	if err != nil {
		return err
	}
	if operr != nil {
		log.Errorf("control | set sock opt err: %s", operr)
		return operr
	}
	return nil
}
