// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

var (
	Stdin  = 0
	Stdout = 1
	Stderr = 2
)

func Errstr(errno int) string {
	if errno < 0 || errno >= int(len(errors)) {
		return "error " + str(errno)
	}
	return errors[errno]
}
