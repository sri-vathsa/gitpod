// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package iws

import (
	"unsafe"

	"golang.org/x/sys/unix"
)

// func syscallMoveMount(fromDirFD int, fromPath string, toDirFD int, toPath string, flags uintptr) error {
// 	fromPathP, err := unix.BytePtrFromString(fromPath)
// 	if err != nil {
// 		return err
// 	}
// 	toPathP, err := unix.BytePtrFromString(toPath)
// 	if err != nil {
// 		return err
// 	}

// 	_, _, errno := unix.Syscall6(unix.SYS_MOVE_MOUNT, uintptr(fromDirFD), uintptr(unsafe.Pointer(fromPathP)), uintptr(toDirFD), uintptr(unsafe.Pointer(toPathP)), flags, 0)
// 	if errno != 0 {
// 		return errno
// 	}

// 	return nil
// }

// const (
// 	// FlagMoveMountFEmptyPath: empty from path permitted: https://elixir.bootlin.com/linux/latest/source/include/uapi/linux/mount.h#L70
// 	flagMoveMountFEmptyPath = 0x00000004
// )

func syscallOpenTree(dfd int, path string, flags uintptr) (fd uintptr, err error) {
	p1, err := unix.BytePtrFromString(path)
	if err != nil {
		return 0, err
	}
	fd, _, errno := unix.Syscall(unix.SYS_OPEN_TREE, uintptr(dfd), uintptr(unsafe.Pointer(p1)), flags)
	if errno != 0 {
		return 0, errno
	}

	return fd, nil
}

const (
	// FlagOpenTreeClone: https://elixir.bootlin.com/linux/latest/source/include/uapi/linux/mount.h#L62
	flagOpenTreeClone = 1
	// FlagAtRecursive: Apply to the entire subtree: https://elixir.bootlin.com/linux/latest/source/include/uapi/linux/fcntl.h#L112
	flagAtRecursive = 0x8000
)
