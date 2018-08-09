package security

import (
	"log"
	"os"
	"syscall"

	"github.com/amy911/srv911/vars"
)

func Chroot() {
	if chroot := vars.Chroot; len(chroot) > 0 {
		if err := syscall.Chroot(chroot); err != nil {
			log.Fatalf("syscall.Chroot: \"%s\"\n", err)
		}
		if err := os.Chdir("/"); err != nil {
			log.Fatalf("os.Chdir: \"%s\"\n", err)
		}
	}
}
