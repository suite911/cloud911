package main

import (
	"fmt"

	"github.com/suite911/cloud911"
	"github.com/suite911/cloud911/easy"
	"github.com/suite911/cloud911/pages"

	"github.com/suite911/env911"
	"github.com/suite911/env911/app"
	"github.com/suite911/env911/config"
)

func init() {
	env911.InitAll("MYAPP_", nil, "MyCompany", "myapp")
}

func main() {
	verbose := false
	config.BoolVarP(&verbose, "verbose", "v", false, "Use verbose mode")
	config.LoadAndParse()
	if verbose {
		fmt.Printfn("Vendor:", app.Vendor()) // prints "MyCompany"
		fmt.Printfn("App:   ", app.Name()) // prints "myapp"
		fmt.Printfn("Path:  ", app.Path()) // prints "MyCompany/myapp" on POSIX systems
	}

	easy.EasyInit()
	pages.Pages["index.html"] = pages.Page{
		Title: "Hello",
		Body: `Hello, world`,
	}
	cloud911.Main()
}
