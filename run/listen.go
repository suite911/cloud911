package run

import (
	"log"

	"github.com/suite911/cloud911/handlers"
	"github.com/suite911/cloud911/vars"

	"github.com/valyala/fasthttp"
)

func Listen() {
	go func() {
		if err := fasthttp.ListenAndServe(vars.HTTP, handlers.HTTP); err != nil {
			log.Fatalln("fasthttp.ListenAndServe: \""+err.Error()+"\"")
		}
	}()
	if err := fasthttp.ListenAndServeTLSEmbed(
		vars.HTTPS,
		vars.CertData,
		vars.KeyData,
		handlers.HTTPS,
	); err != nil {
		log.Fatalln("fasthttp.ListenAndServeTLSEmbed: \""+err.Error()+"\"")
	}
}
