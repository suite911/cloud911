package main

// go:generate slurp911 -o g_www.go -n www www

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

	if raw, ok := www["favicon.ico"]; ok {
		pages.Pages["favicon.ico"] = pages.Page{
			Raw: raw,
		}
	}

	pages.Pages["index.html"] = pages.Page{
		Title: "My App",
		Body: string(www["index.htm"]),
		CSS: string(www["index.css"]),
	}

	cloud911.Main(exampleCallback)
}

func exampleCallback() error {
	if verbose {
		fmt.Println("Vendor:", app.Vendor()) // prints "MyCompany"
		fmt.Println("App:   ", app.Name()) // prints "myapp"
		fmt.Println("Path:  ", app.Path()) // prints "MyCompany/myapp" on POSIX systems
	}
	return nil
}

var www = map[string][]byte{}
