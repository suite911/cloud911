package main

import (
	"github.com/amy911/snek911/snek"

	"github.com/amy911/srv911/run"
	"github.com/amy911/srv911/vars"

	"github.com/spf13/cobra"
)

func main() {
	snek.InitRoot = func(cmd *cobra.Command) {
		cmd.PersistentFlags().StringVar(&vars.AddrHttp, "http", "", "Address on which to listen to incoming HTTP traffic")
		cmd.PersistentFlags().StringVar(&vars.AddrHttps, "https", "", "Address on which to listen to incoming HTTPS traffic")
		cmd.PersistentFlags().StringVar(&vars.Chroot, "chroot", "", "Path to which to chroot(2)")
		cmd.PersistentFlags().StringVar(&vars.CertPath, "cert", "", "Path of TLS certificate file")
		cmd.PersistentFlags().StringVar(&vars.KeyPath, "key", "", "Path of TLS key file")
		snek.Bind("http", "https", "chroot", "cert", "key")
		cmd.Short = "Listen and serve"
		cmd.Long = `Listen and serve`
		cmd.Run = run.Listen
	}
	snek.Main()
}
