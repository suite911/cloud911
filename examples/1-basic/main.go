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

	pages.Pages["index.html"] = pages.Page{
		Title: "My App",
		Body: www["index.htm"],
		CSS: www["index.css"],
	}

	pages.Pages["about"] = pages.Page{
		Title: "My App - About",
		Body: www["about.htm"],
		CSS: www["about.css"],
	}

	pages.Pages["register"] = pages.Page{
		Title: "My App - Register",
		Body: www["register.htm"],
		CSS: www["register.css"],
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
