package main

import (
	"fmt"

	"github.com/suite911/cloud911"
	"github.com/suite911/cloud911/pages"

	"github.com/suite911/env911"
	"github.com/suite911/env911/app"
	"github.com/suite911/env911/config"
)

func init() {
	env911.InitAll("MYAPP_", nil, "MyCompany", "myapp")
}

var verbose = false

func main() {
	flagSet := config.FlagSet()
	flagSet.BoolVarP(&verbose, "verbose", "v", false, "Use verbose mode")

	pages.Pages["index.html"] = pages.Page{
		Title: "Hello",
		Body: `Hello, world`,
	}

	cloud911.Main(exampleCallback)
}

func exampleCallback() error {
	if verbose {
		fmt.Println("Vendor:", app.Vendor()) // prints "MyCompany"
		fmt.Println("App:   ", app.Name()) // prints "myapp"
		fmt.Println("Path:  ", app.Path()) // prints "MyCompany/myapp" on POSIX systems
	}
}
