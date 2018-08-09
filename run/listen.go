package run

import (
	"log"

	"github.com/amy911/srv911/handlers"
	"github.com/amy911/srv911/security"
	"github.com/amy911/srv911/vars"

	"github.com/spf13/cobra"

	"github.com/valyala/fasthttp"
)

func Listen(cmd *cobra.Command, args []string) {
	security.Chroot()
	if err := fasthttp.ListenAndServe(vars.AddrHttp, handlers.Root); err != nil {
		log.Fatalln("fasthttp.ListenAndServe: \""+err.Error()+"\"")
	}
}
