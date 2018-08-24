package run

import (
	"log"
	"net"

	"github.com/suite911/cloud911/handlers"
	"github.com/suite911/cloud911/pages"
	"github.com/suite911/cloud911/shells"
	"github.com/suite911/cloud911/vars"

	"github.com/valyala/fasthttp"
)

func Listen(http, https net.Listener) error {
	if err := pages.Compile(shells.Amy); err != nil {
		return err
	}
	go func() {
		if err := fasthttp.Serve(http, handlers.HTTP); err != nil {
			log.Fatalln("fasthttp.ListenAndServe: \""+err.Error()+"\"")
		}
	}()
	if err := fasthttp.ServeTLSEmbed(
		https,
		vars.Pass.TLSCertData,
		vars.Pass.TLSKeyData,
		handlers.HTTPS,
	); err != nil {
		log.Fatalln("fasthttp.ListenAndServeTLSEmbed: \""+err.Error()+"\"")
	}
	return nil
}
