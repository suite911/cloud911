package cloud911

import (
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
		"The following are recognized for OPTIONS:\n",
	)

	if len(os.Args) < 2 {
		stdin := os.Stdin.Fd()
		if isatty.IsTerminal(stdin) || isatty.IsCygwinTerminal(stdin) {
			// Run without args from a terminal: Usage
			config.Usage()
			os.Exit(0)
		}
		// Child
		return child(fns)
	}
	// Parent
	return parent()
}
