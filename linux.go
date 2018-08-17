// +build linux

package cloud911

import (
	"os/exec"

	"golang.org/x/sys/unix"
)

func ApplyLinuxCloneFlags(cmd *exec.Cmd) {
	if cmd.SysProcAttr == nil {
		cmd.SysProcAttr = new(unix.SysProcAttr)
	}
	cmd.SysProcAttr.Cloneflags |= unix.CLONE_NEWNS | unix.CLONE_NEWPID | unix.CLONE_NEWUTS
}
