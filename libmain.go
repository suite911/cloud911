package cloud911

import (
	"io/ioutil"
	"log" // TODO
	"net"
	"os"
	"strconv"

	"github.com/suite911/cloud911/database"
	"github.com/suite911/cloud911/droppriv"
	"github.com/suite911/cloud911/run"
	"github.com/suite911/cloud911/vars"

	"github.com/suite911/env911"
	"github.com/suite911/env911/config"
	"github.com/suite911/term911/vt"

	"github.com/pkg/errors"
	"golang.org/x/sys/unix"
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
		"Usage: " + os.Args[0] + " " +
		vt.SafeS("[") + vt.SafeU("OPTIONS") + vt.SafeS("]") + "\n" +
		"Usage: " + vt.SafeB("sudo") + " " + os.Args[0] + " " +
		vt.SafeS("[") + vt.SafeU("OPTIONS") + vt.SafeS("]") + " " +
		vt.SafeU("UID") + " " + vt.SafeU("GID") + "\n" +
		"\n" +
		"The following are recognized for " + vt.SafeU("OPTIONS") + ":\n",
	)

	uid, gid := os.Getuid(), os.Getgid()
	args := flagSet.Args()
	switch {
	case uid < 0 || gid < 0: // Windows
		log.Fatalln("Operating system not supported.")
	case uid > 0 && gid > 0:
		if len(args) != 0 {
			flagSet.Usage()
			os.Exit(1)
		}
		if len(os.Args) < 2 {
			flagSet.Usage()
			os.Exit(0)
		}
	case uid == 0:
		if len(args) != 2 {
			flagSet.Usage()
			os.Exit(1)
		}
	default:
		flagSet.Usage()
		os.Exit(1)
	}

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
			return errors.Wrap(err, "os.Chdir(chroot)")
		}
		if err := unix.Chroot(chroot); err != nil {
			return errors.Wrap(err, "os.Chroot(chroot)")
		}
		if err := os.Chdir("/"); err != nil {
			return errors.Wrap(err, "os.Chdir(\"/\")")
		}
	}
	http, err := net.Listen("tcp4", vars.Pass.HTTP)
	if err != nil {
		return err
	}
	https, err := net.Listen("tcp4", vars.Pass.HTTPS)
	if err != nil {
		return err
	}

	if len(args) >= 2 {
		if vars.UID, err = strconv.Atoi(args[0]); err != nil {
			return err
		}
		if vars.GID, err = strconv.Atoi(args[1]); err != nil {
			return err
		}
	}

	// Drop privileges

	if err := droppriv.Drop(); err != nil {
		return err
	}

	for _, fn := range fns {
		if err := fn(); err != nil {
			return err
		}
	}

	if err := database.Create(); err != nil {
		return err
	}

	return run.Listen(http, https)
}

func loadTLSCert() error {
	var err error
	certPath, keyPath := vars.CertPath, vars.KeyPath
	if vars.Pass.TLSCertData, err = ioutil.ReadFile(certPath); err != nil {
		return tlsReadFileError(certPath, keyPath, errors.Wrap(err, "ioutil.ReadFile(certPath)"))
	}
	if vars.Pass.TLSKeyData, err = ioutil.ReadFile(keyPath); err != nil {
		return tlsReadFileError(certPath, keyPath, errors.Wrap(err, "ioutil.ReadFile(keyPath)"))
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
