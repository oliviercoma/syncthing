// Copyright (C) 2015 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// +build !windows

package osutil

import (
	"os"
	"syscall"
)

func Uid(info os.FileInfo) int {
	return info.Sys.(*syscall.Stat_t).Uid;
}

func Gid(info os.FileInfo) int {
	return info.Sys.(*syscall.Stat_t).Gid;
}
