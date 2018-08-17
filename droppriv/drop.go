package droppriv

import (
	"errors"

	"github.com/suite911/cloud911/vars"

	pkgErrors "github.com/pkg/errors"
	"golang.org/x/sys/unix"
)

func Drop() error {
	uid, gid := vars.UID, vars.GID
	if uid < 1 || gid < 1 {
		return pkgErrors.WithStack(errors.New("Bad UID or GID!"))
	}
	if err := syscall1(unix.SYS_SETUID, uid); err != nil {
		return pkgErrors.Wrap(err, "SYS_SETUID")
	}
	if err := syscall1(unix.SYS_SETGID, gid); err != nil {
		return pkgErrors.Wrap(err, "SYS_SETGID")
	}
	if err := syscall1(unix.SYS_SETEUID, uid); err != nil {
		return pkgErrors.Wrap(err, "SYS_SETEUID")
	}
	if err := syscall1(unix.SYS_SETEGID, gid); err != nil {
		return pkgErrors.Wrap(err, "SYS_SETEGID")
	}
	newUID, newGID := os.Getuid(), os.Getgid()
	if newUID == 0 || newGID == 0 {
		return pkgErrors.WithStack(errors.New("Unable to drop privileges!"))
	}
	return nil
}

func syscall1(trap uintptr, arg int) {
	var err error
	_, _, en := unix.Syscall(trap, uintptr(arg), 0, 0)
	if en != 0 {
		err = en
	}
	return err
}
