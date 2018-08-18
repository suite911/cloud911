package droppriv

import (
	"errors"
	"os"

	"github.com/suite911/cloud911/vars"

	pkgErrors "github.com/pkg/errors"
	"golang.org/x/sys/unix"
)

func Drop() error {
	uid, gid := vars.UID, vars.GID
	if uid < 1 || gid < 1 {
		return pkgErrors.WithStack(errors.New("Bad UID or GID!"))
	}
	if os.Getuid() == 0 {
		if err := syscall2(SYS_SETREUID, uid); err != nil {
			return pkgErrors.Wrap(err, "SYS_SETREUID")
		}
		if os.Getuid() == 0 {
			return pkgErrors.WithStack(errors.New("Unable to drop uid 0!"))
		}
	}
	if os.Getgid() == 0 {
		if err := syscall2(SYS_SETREGID, gid); err != nil {
			return pkgErrors.Wrap(err, "SYS_SETREGID")
		}
		if os.Getgid() == 0 {
			return pkgErrors.WithStack(errors.New("Unable to drop gid 0!"))
		}
	}
	return nil
}

func syscall2(trap uintptr, arg int) error {
	var err error
	_, _, en := unix.Syscall(trap, uintptr(arg), uintptr(arg), uintptr(arg))
	if en != 0 {
		err = en
	}
	return err
}
