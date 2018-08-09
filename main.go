package main

import (
	"github.com/amy911/srv911/cmd"
	_ "github.com/amy911/srv911/secret"
)

func main() {
	cmd.Execute()
}
