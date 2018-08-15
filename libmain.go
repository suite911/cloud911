package cloud911

import (
	"errors"
	"io"
	"os"
	"os/exec"

	"github.com/amy911/cloud911/run"
	"github.com/amy911/cloud911/vars"

	"github.com/amy911/env911"
	"github.com/amy911/env911/config"
	"github.com/amy911/term911/vt"

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
	if pchroot == nil {
		// This can happen if the user's custom FlagSet instance is broken
		panic("Something is wrong with the custom github.com/amy911/flag911/flag.FlagSet you used with github.com/amy911/env911[/config]")
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
	str := "Hello, world"
	n, err := io.WriteString(stdin, str)
	if err != nil {
		return err
	}
	if n != len(str) { // just in case
		return errors.New("Write error")
	}
	return child.Run()
}

func child(fns []func() error) error {
	for _, fn := range fns {
		if err := fn(); err != nil {
			return err
		}
	}
	run.Listen()
}
