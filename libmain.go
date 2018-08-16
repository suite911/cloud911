package cloud911

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log" // TODO
	"os"
	"os/exec"

	"github.com/suite911/cloud911/run"
	"github.com/suite911/cloud911/vars"

	"github.com/suite911/env911"
	"github.com/suite911/env911/config"
	"github.com/suite911/term911/vt"

	"github.com/mattn/go-isatty"
)

func Main(fns ...func() error) error {
	if !env911.IsInitAll() {
		panic("Please initialize env911 first!")
	}
	flagSet := config.FlagSet()
	flagSet.StringVar(&vars.Pass.HTTP, "http", "", "Address on which to listen to incoming HTTP traffic")
	flagSet.StringVar(&vars.Pass.HTTPS, "https", "", "Address on which to listen to incoming HTTPS traffic")
	pchroot := flagSet.String("chroot", "", "Path to which to chroot(2)")
	flagSet.StringVar(&vars.CertPath, "cert", "", "Path of TLS certificate file")
	flagSet.StringVar(&vars.KeyPath, "key", "", "Path of TLS key file")
	config.LoadAndParse()

	flagSet.SetUsageHeader(
		os.Args[0] + " " + vt.U("VERB") + " " + vt.U("OPTIONS") + vt.NewLine +
		"The following are recognized for VERB:\n" +
		"    " + vt.U("help") + "  \t: Print this help text and exit.\n" +
		"    " + vt.U("listen") + "\t: Listen and serve.\n" +
		"The following are recognized for OPTIONS:\n"
	)

	if len(os.Args) < 2 {
		stdin := os.Stdin.Fd()
		if isatty.IsTerminal(stdin) || isatty.IsCygwinTerminal(stdin) {
			config.Usage()
			os.Exit(0)
		}
		return child(fns)
	}
	// Parent
	return parent()
}

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

func child(fns []func() error) error {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, &vars.Pass); err != nil {
		return err
	}
	for _, fn := range fns {
		if err := fn(); err != nil {
			return err
		}
	}
	run.Listen()
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
