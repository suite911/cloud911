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

	pages.Pages[""] = pages.Page{
		FavIcon: favIcon,
		Footer: string(www["footer.htm"]),
		Shell: shells.Amy,
		Title: "My App",
		TopNav: topNav,
	}

	pages.Pages["/about"] = pages.Page{
		FavIcon: favIcon,
		Footer: string(www["footer.htm"]),
		Shell: shells.Amy,
		Title: "My App - About",
		TopNav: topNav,
	}

	pages.Pages["/eula"] = pages.Page{
		FavIcon: favIcon,
		Footer: string(www["eula.htm"]),
		Shell: shells.Amy,
		Title: "My App - End User License Agreement (EULA)",
		TopNav: topNav,
	}

	pages.Pages["/privacy"] = pages.Page{
		Content: string(www["privacy.htm"]),
		FavIcon: favIcon,
		Footer: string(www["footer.htm"]),
		Shell: shells.Amy,
		Title: "My App - Privacy Policy",
		TopNav: topNav,
	}

	pages.Pages["/register"] = pages.Page{
		FavIcon: favIcon,
		Footer: string(www["footer.htm"]),
		Shell: shells.Amy,
		Title: "My App - Register",
		TopNav: topNav,
	}

	pages.Pages["/robots.txt"] = pages.Page{
		ContentType: "text/plain; charset=utf8",
		Raw: www["robots.txt"],
	}

	pages.Pages["/terms"] = pages.Page{
		Content: string(www["terms.htm"]),
		FavIcon: favIcon,
		Footer: string(www["footer.htm"]),
		Shell: shells.Amy,
		Title: "My App - Terms of Use",
		TopNav: topNav,
	}

	cloud911.Main()
}

var www = map[string][]byte{}
