package cloud911

import (
	"github.com/amy911/cloud911/run"
	"github.com/amy911/cloud911/vars"

	"github.com/amy911/env911"
	"github.com/amy911/env911/config"
	"github.com/amy911/env911/vt"
)

func Main(fns ...func() error) error {
	if !env911.IsInitAll() {
		panic("Please initialize env911 first!")
	}
	config.StringVar(&vars.AddrHttp, "http", "", "Address on which to listen to incoming HTTP traffic")
	config.StringVar(&vars.AddrHttps, "https", "", "Address on which to listen to incoming HTTPS traffic")
	pchroot := config.String("chroot", "", "Path to which to chroot(2)")
	config.StringVar(&vars.CertPath, "cert", "", "Path of TLS certificate file")
	config.StringVar(&vars.KeyPath, "key", "", "Path of TLS key file")
	config.LoadAndParse()

	config.SetUsageHeader(
		os.Args[0] + " " + vt.U("VERB") + " " + vt.U("OPTIONS") + vt.NewLine +
		"The following are recognized for VERB:\n" +
		"    " + vt.U("help") + "  \t: Print this help text and exit.\n" +
		"    " + vt.U("listen") + "\t: Listen and serve.\n" +
		"The following are recognized for OPTIONS:\n"
	)

	if len(os.Args) < 2 {
		config.Usage()
		os.Exit(0)
	}
	switch os.Args[1] {
	case "help":
		config.Usage()
		os.Exit(0)
	case "listen":
		// Parent
		if pchroot == nil {
			// This can happen if the user's custom Flagger instance is broken
			panic("Something is wrong with the Flagger you used with github.com/amy911/env911[/config]")
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
		return parent()
	case "!":
		// Child
		return child()
	default:
		config.Usage()
		os.Exit(1)
	}

	if len(os.Args) > 1 && os.Args[1] == "!" {
		return child(fns)
	}
	if len(*pchroot) > 0 {
		return parent()
	}
	return child(fns)
}

func parent() error {
	child := exec.Command(os.Executable(), append([]string{"!"}, os.Args[1:]))
	child.SysProcAttr = &syscall.SysProcAttr{
		CloneFlags: syscall.CLONE_NEWNS | syscall.CLONE_NEWPID | syscall.CLONE_NEWUTS,
	}
	child.Stdin = os.Stdin
	child.Stdout = os.Stdout
	child.Stderr = os.Stderr
	return child.Run()
}

func child(fns []func() error) error {
	if err := syscall.Mount("rootfs", "rootfs", "", syscall.MS_BIND, ""); err != nil {
		return err
	}
	if err := os.MkdirAll("rootfs/old", 0700); err != nil {
		return err
	}
	if err := syscall.PivotRoot("rootfs", "rootfs/old"); err != nil {
		return err
	}
	if err := os.Chdir("/"); err != nil {
		return err
	}
	if err := os.Chdir(".."); err == nil {
		return errors.New("Trivial escape from chroot possible!")
	}
	return runChild(fns)
}

func runChild(fns []func() error) error {
	for _, fn := range fns {
		if err := fn(); err != nil {
			return err
		}
	}
	run.Listen()
}
