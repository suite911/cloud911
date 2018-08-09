package run

import (
	"crypto/tls"
	"log"

	"github.com/amy911/amy911/onfail"
	"github.com/amy911/amy911/security"

	"github.com/amy911/srv911/handlers"
	"github.com/amy911/srv911/vars"

	"github.com/spf13/cobra"

	"github.com/valyala/fasthttp"
)

func Listen(cmd *cobra.Command, args []string) {
	if chroot := vars.Chroot; len(chroot) > 0 {
		if err := security.Chroot(chroot, onfail.Fatal); err != nil {
			panic(err) // just in case
		}
	}
	go func(cmd *cobra.Command, args []string) {
		if err := fasthttp.ListenAndServe(vars.AddrHttp, handlers.Http); err != nil {
			log.Fatalln("fasthttp.ListenAndServe: \""+err.Error()+"\"")
		}
	}(cmd, args)
	certPath, keyPath := vars.CertPath, vars.KeyPath
	if err := fasthttp.ListenAndServeTLS(
		vars.AddrHttps,
		certPath,
		keyPath,
		handlers.Https,
	); err != nil {
		if _, err := tls.LoadX509KeyPair(certPath, keyPath); err != nil {
			log.Printf(
				"You need a TLS certificate file and a TLS key file.  "+
				"By default, these are called \"cert.pem\" and \"key.pem\", respectively.  "+
				"The paths as configured are %q and %q.", certPath, keyPath)
		}
		log.Fatalln("fasthttp.ListenAndServeTLS: \""+err.Error()+"\"")
	}
}
