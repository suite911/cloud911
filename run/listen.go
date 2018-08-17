package run

import (
	"log"

	"github.com/suite911/cloud911/handlers"
	"github.com/suite911/cloud911/pages"
	"github.com/suite911/cloud911/shells"
	"github.com/suite911/cloud911/vars"

	"github.com/valyala/fasthttp"
)

func Listen() error {
	if err := pages.PreparePageBytes(shells.Basic); err != nil {
		return err
	}
	go func() {
		if err := fasthttp.ListenAndServe(vars.Pass.HTTP, handlers.HTTP); err != nil {
			log.Fatalln("fasthttp.ListenAndServe: \""+err.Error()+"\"")
		}
	}()
	if err := fasthttp.ListenAndServeTLSEmbed(
		vars.Pass.HTTPS,
		vars.Pass.TLSCertData,
		vars.Pass.TLSKeyData,
		handlers.HTTPS,
	); err != nil {
		log.Fatalln("fasthttp.ListenAndServeTLSEmbed: \""+err.Error()+"\"")
	}
	return nil
}
