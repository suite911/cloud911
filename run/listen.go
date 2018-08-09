package run

import (
	"os"
	"syscall"

	"github.com/amy911/srv911/handlers"
	"github.com/amy911/srv911/vars"

	"github.com/spf13/cobra"

	"github.com/valyala/fasthttp"
)

func Listen(cmd *cobra.Command, args []string) {
	if chroot := vars.Chroot; len(chroot) > 0 {
		syscall.Chroot(chroot)
	}
	os.Chdir(".")
	if err := fasthttp.ListenAndServe(vars.AddrHttp, handlers.Root); err != nil {
		log.Fatalln("fasthttp.ListenAndServe: \""+err.Error()+"\"")
	}
}
