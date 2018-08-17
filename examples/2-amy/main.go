package main

// go:generate slurp911 -o g_www.go -n www www

import (
	"fmt"

	"github.com/suite911/cloud911"
	"github.com/suite911/cloud911/pages"
	"github.com/suite911/cloud911/shells"

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

	topNav := make(map[string]string)
	topNav["/"] = "Top Page"
	topNav["/about"] = "About"
	topNav["/register"] = "Register"

	var favIcon string
	if raw, ok := www["favicon.ico"]; ok {
		pages.Pages["/favicon.ico"] = pages.Page{
			Raw: raw,
		}
		favIcon = "favicon.ico"
	}

	pages.Pages["/index.html"] = pages.Page{
		FavIcon: favIcon,
		Shell: shells.Amy,
		Title: "My App",
		TopNav: topNav,
	}

	pages.Pages["/about"] = pages.Page{
		FavIcon: favIcon,
		Shell: shells.Amy,
		Title: "My App - About",
		TopNav: topNav,
	}

	pages.Pages["/register"] = pages.Page{
		FavIcon: favIcon,
		Shell: shells.Amy,
		Title: "My App - Register",
		TopNav: topNav,
	}

	cloud911.Main()
}

var www = map[string][]byte{}
