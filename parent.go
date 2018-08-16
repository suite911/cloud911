package cloud911

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log" // TODO
	"os"
	"os/exec"
	"syscall"
)

func parent() error {
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
			return err
		}
		if err := os.Chroot(chroot); err != nil {
			return err
		}
		if err := os.Chdir("/"); err != nil {
			return err
		}
		if err := os.Chdir(".."); err == nil {
			return errors.New("Trivial escape from chroot possible!")
		}
	}
	child := exec.Command(os.Executable())
	child.SysProcAttr = &syscall.SysProcAttr{
		CloneFlags: syscall.CLONE_NEWNS | syscall.CLONE_NEWPID | syscall.CLONE_NEWUTS,
	}
	child.Stdout = os.Stdout
	child.Stderr = os.Stderr
	stdin, err := child.StdinPipe()
	if err != nil {
		return err
	}
	b, err := json.Marshal(vars.Pass)
	if err != nil {
		return err
	}
	n, err := stdin.Write(b)
	if err != nil {
		return err
	}
	if n != len(b) { // just in case
		return errors.New("Write error")
	}
	return child.Run()
}

func loadTLSCert() error {
	certPath, keyPath := vars.CertPath, vars.KeyPath
	certData, err := ioutil.ReadFile(certPath)
	if err != nil {
		return tlsReadFileError(certPath, keyPath, err)
	}
	keyData, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return tlsReadFileError(certPath, keyPath, err)
	}
	return nil
}

func tlsReadFileError(certPath, keyPath string, err error) error {
	log.Printf(
		"You need a TLS certificate file and a TLS key file.  "+
		"By default, these are called \"cert.pem\" and \"key.pem\", respectively.  "+
		"The paths as configured are %q and %q.", certPath, keyPath)
	log.Fatalf("ioutil.ReadFile: %q\n", err)
	return err
}
