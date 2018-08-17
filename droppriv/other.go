// +build !linux !freebsd

package droppriv

import "golang.org/x/sys/unix"

const (
	SYS_SETUID = unix.SYS_SETUID
	SYS_SETGID = unix.SYS_SETGID
	SYS_SETEUID = unix.SYS_SETEUID
	SYS_SETEGID = unix.SYS_SETEGID
)

func LinuxDrop() error {
	return nil
}
