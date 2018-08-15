package run

import (
	"io/ioutil"
	"log"

	"github.com/suite911/error911/onfail"
	"github.com/suite911/amy911/security"

	"github.com/suite911/cloud911/handlers"
	"github.com/suite911/cloud911/vars"

	"github.com/spf13/cobra"

	"github.com/valyala/fasthttp"
)

func Listen(cmd *cobra.Command, args []string) {
	certPath, keyPath := vars.CertPath, vars.KeyPath
	certData, err := ioutil.ReadFile(certPath)
	if err != nil {
		tlsReadFileError(certPath, keyPath, err)
	}
	keyData, err := ioutil.ReadFile(keyPath)
	if err != nil {
		tlsReadFileError(certPath, keyPath, err)
	}
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
	if err := fasthttp.ListenAndServeTLSEmbed(
		vars.AddrHttps,
		certData,
		keyData,
		handlers.Https,
	); err != nil {
		log.Fatalln("fasthttp.ListenAndServeTLS: \""+err.Error()+"\"")
	}
}

func tlsReadFileError(certPath, keyPath string, err error) {
	log.Printf(
		"You need a TLS certificate file and a TLS key file.  "+
		"By default, these are called \"cert.pem\" and \"key.pem\", respectively.  "+
		"The paths as configured are %q and %q.", certPath, keyPath)
	log.Fatalf("ioutil.ReadFile: %q\n", err)
}
