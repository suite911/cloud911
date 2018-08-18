// +build freebsd

package droppriv

const (
	SYS_setuid = uintptr(23)
	SYS_setgid = uintptr(181)
	SYS_seteuid = uintptr(183)
	SYS_setegid = uintptr(182)
)

func LinuxDrop() error {
	return nil
}
