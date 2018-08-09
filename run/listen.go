package run

import (
	"os"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/valyala/fasthttp"
)

func Listen(cmd *cobra.Command, args []string) {
	if chroot := vars.Chroot; len(chroot) > 0 {
		syscall.Chroot(chroot)
	}
	os.Chdir(".")
}
