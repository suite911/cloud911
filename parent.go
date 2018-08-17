package cloud911

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log" // TODO
	"os"
	"os/exec"

	"github.com/suite911/cloud911/vars"

	pkgErrors "github.com/pkg/errors"
	"golang.org/x/sys/unix"
)

func parent(pchroot *string) error {
	if err := loadTLSCert(); err != nil {
		return err
	}
	if pchroot == nil {
		// This can happen if the user's custom FlagSet instance is broken
		panic("Something is wrong with the custom github.com/suite911/flag911/flag.FlagSet you used with github.com/suite911/env911[/config]")
	}
	chroot := *pchroot
	if len(chroot) > 0 {
		if err := os.Chdir(chroot); err != nil {
			return pkgErrors.Wrap(err, "os.Chdir(chroot)")
		}
		if err := unix.Chroot(chroot); err != nil {
			return pkgErrors.Wrap(err, "os.Chroot(chroot)")
		}
		if err := os.Chdir("/"); err != nil {
			return pkgErrors.Wrap(err, "os.Chdir(\"/\")")
		}
		if err := os.Chdir(".."); err == nil {
			return errors.New("Trivial escape from chroot possible!")
		}
	}
	self, err := os.Executable()
	if err != nil {
		return pkgErrors.Wrap(err, "os.Executable")
	}
	child := exec.Command(self)
	ApplyLinuxCloneFlags(child)
	child.Stdout = os.Stdout
	child.Stderr = os.Stderr
	stdin, err := child.StdinPipe()
	if err != nil {
		return pkgErrors.Wrap(err, "(child *exec.Cmd).StdinPipe()")
	}
	b, err := json.Marshal(vars.Pass)
	if err != nil {
		return pkgErrors.Wrap(err, "json.Marshall")
	}
	n, err := stdin.Write(b)
	if err != nil {
		return pkgErrors.Wrap(err, "(stdin io.WriteCloser).Write")
	}
	if n != len(b) { // just in case
		return errors.New("Write error")
	}
	return child.Run()
}

func loadTLSCert() error {
	var err error
	certPath, keyPath := vars.CertPath, vars.KeyPath
	if vars.Pass.TLSCertData, err = ioutil.ReadFile(certPath); err != nil {
		return tlsReadFileError(certPath, keyPath, pkgErrors.Wrap(err, "ioutil.ReadFile(certPath)"))
	}
	if vars.Pass.TLSKeyData, err = ioutil.ReadFile(keyPath); err != nil {
		return tlsReadFileError(certPath, keyPath, pkgErrors.Wrap(err, "ioutil.ReadFile(keyPath)"))
	}
	return nil
}

func tlsReadFileError(certPath, keyPath string, err error) error { // TODO
	log.Printf(
		"You need a TLS certificate file and a TLS key file.  "+
		"By default, these are called \"cert.pem\" and \"key.pem\", respectively.  "+
		"The paths as configured are %q and %q.", certPath, keyPath)
	log.Fatalf("ioutil.ReadFile: %q\n", err)
	return err
}
