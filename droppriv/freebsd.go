// +build freebsd

package droppriv

import "golang.org/x/sys/unix"

const (
	SYS_SETUID = unix.SYS_SETUID
	SYS_SETGID = unix.SYS_SETGID
	SYS_SETEUID = uintptr(183)
	SYS_SETEGID = uintptr(182)
)

func LinuxDrop() error {
	return nil
}
