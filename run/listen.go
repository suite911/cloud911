package run

import (
	"log"

	"github.com/suite911/cloud911/handlers"
	"github.com/suite911/cloud911/vars"

	"github.com/valyala/fasthttp"
)

func Listen(cmd *cobra.Command, args []string) {
	go func(cmd *cobra.Command, args []string) {
		if err := fasthttp.ListenAndServe(vars.HTTP, handlers.HTTP); err != nil {
			log.Fatalln("fasthttp.ListenAndServe: \""+err.Error()+"\"")
		}
	}(cmd, args)
	if err := fasthttp.ListenAndServeTLSEmbed(
		vars.HTTPS,
		certData,
		keyData,
		handlers.HTTPS,
	); err != nil {
		log.Fatalln("fasthttp.ListenAndServeTLS: \""+err.Error()+"\"")
	}
}
