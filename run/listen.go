package run

import (
	"log"

	"github.com/amy911/srv911/handlers"
	"github.com/amy911/srv911/util/security"
	"github.com/amy911/srv911/vars"

	"github.com/spf13/cobra"

	"github.com/valyala/fasthttp"
)

func Listen(cmd *cobra.Command, args []string) {
	security.Chroot()
	go func(cmd *cobra.Command, args []string) {
		if err := fasthttp.ListenAndServe(vars.AddrHttp, handlers.Http); err != nil {
			log.Fatalln("fasthttp.ListenAndServe: \""+err.Error()+"\"")
		}
	}(cmd, args)
	go func(cmd *cobra.Command, args []string) {
		if err := fasthttp.ListenAndServeTLS(vars.AddrHttps, handlers.Https); err != nil {
			log.Fatalln("fasthttp.ListenAndServeTLS: \""+err.Error()+"\"")
		}
	}(cmd, args)
}
