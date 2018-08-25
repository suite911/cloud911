package main

//go:generate slurp911 -o=g_www.go -n=www www

import (
	"github.com/suite911/cloud911"
	"github.com/suite911/cloud911/pages"
	"github.com/suite911/cloud911/shells"
	"github.com/suite911/cloud911/vars"

	"github.com/suite911/env911"
)

func init() {
	env911.InitAll("MYAPP_", nil, "MyCompany", "myapp")
}

var verbose = false

func main() {
	vars.CaptchaSiteKey = "6LeB-WsUAAAAANf-vw8q2NtTybrb85G3HyEBXyxO"

	topNav := map[string]string{
		"/": "&#x1f3e0;",
		"/download": "Download",
		"/gallery": "Gallery",
		"/register": "Register",
		"/store": "Store",
	}

	pages.Pages["404"] = &pages.Page{
		Content: string(www["/404.htm"]),
		ContentSubTitle: "That's an HTTP 404 right there.",
		ContentTitle: "Not Found",
		PageTitle: "My App - Not Found",
	}

	pages.Pages["/"] = &pages.Page{
		Content: string(www["/index.htm"]),
		ContentTitle: "My App",
		PageTitle: "My App",
	}

	pages.Pages["/about"] = &pages.Page{
		Content: string(www["/about.htm"]),
		ContentTitle: "About",
		PageTitle: "My App - About",
	}

	pages.Pages["/cookies"] = &pages.Page{
		Content: string(www["/cookies.htm"]),
		ContentTitle: "Cookie Policy",
		PageTitle: "My App - Cookie Policy",
	}

	pages.Pages["/download"] = &pages.Page{
		Content: string(www["/download.htm"]),
		ContentTitle: "Downloads",
		PageTitle: "My App - Download",
	}

	pages.Pages["/downloads"] = &pages.Page{
		Redirect301: []byte("/download"),
	}

	pages.Pages["/eula"] = &pages.Page{
		Content: string(www["/eula.htm"]),
		ContentTitle: "End User License Agreement (EULA)",
		PageTitle: "My App - End User License Agreement (EULA)",
	}

	pages.Pages["/gallery"] = &pages.Page{
		Content: string(www["/gallery.htm"]),
		ContentTitle: "Gallery",
		PageTitle: "My App - Gallery",
	}

	pages.Pages["/privacy"] = &pages.Page{
		Content: string(www["/privacy.htm"]),
		ContentTitle: "Privacy Policy",
		PageTitle: "My App - Privacy Policy",
	}

	pages.Pages["/register"] = &pages.Page{
		Content: string(www["/register.htm"]),
		ContentTitle: "Register",
		ContentSubTitle: "Create a new account.",
		Form: true,
		NoScript: "Hello, fellow NoScript user!  This is awkward but could you pretty please whitelist " +
			"my registration page?  You see, the Captcha gods are picky and like to use JavaScript " +
			"to do ...well whatever Captcha gods do with JavaScript.  So anyway it would just make " +
			"things a whole lot easier if you could just whitelist my registration page, kthx.",
		PageTitle: "My App - Register",
	}
	vars.Pass.AlreadyRegistered = "/download#already-registered"
	vars.Pass.Registered = "/download#registered"

	pages.Pages["/robots.txt"] = &pages.Page{
		ContentType: "text/plain; charset=utf8",
		Raw: www["/robots.txt"],
	}

	pages.Pages["/store"] = &pages.Page{
		Content: string(www["/store.htm"]),
		ContentTitle: "Store",
		PageTitle: "My App - Store",
	}

	pages.Pages["/terms"] = &pages.Page{
		Content: string(www["/terms.htm"]),
		ContentTitle: "Terms of Service",
		PageTitle: "My App - Terms of Service",
	}

	var favIcon string
	if raw, ok := www["/favicon.ico"]; ok {
		pages.Pages["/favicon.ico"] = &pages.Page{
			Raw: raw,
		}
		favIcon = "favicon.ico"
	}

	for _, k := range []string{
		"404",
		"/",
		"/about",
		"/cookies",
		"/download",
		"/eula",
		"/gallery",
		"/privacy",
		"/register",
		"/store",
		"/terms",
	} {
		if p, ok := pages.Pages[k]; ok {
			if len(favIcon) > 0 {
				if len(p.FavIcon) < 1 {
					p.FavIcon = favIcon
				}
			}
			if len(p.Footer) < 1 {
				p.Footer = string(www["/footer.htm"])
			}
			if len(p.GoogleFonts) < 1 {
				p.GoogleFonts = "Noto+Sans|Source+Code+Pro"
			}
			if p.Shell == nil {
				p.Shell = shells.Amy
			}
			if len(p.TopNav) < 1 {
				p.TopNav = topNav
			}
		}
	}

	if err := cloud911.Main(); err != nil {
		panic(err)
	}
}

var www = make(map[string][]byte)
