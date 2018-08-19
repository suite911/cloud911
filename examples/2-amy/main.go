package main

//go:generate slurp911 -o=g_www.go -n=www www

import (
	"github.com/suite911/cloud911"
	"github.com/suite911/cloud911/pages"
	"github.com/suite911/cloud911/shells"

	"github.com/suite911/env911"
)

func init() {
	env911.InitAll("MYAPP_", nil, "MyCompany", "myapp")
}

var verbose = false

func main() {
	pages.Pages[""] = &pages.Page{
		Title: "My App",
	}

	pages.Pages["404"] = &pages.Page{
		Body: string(www["/404.htm"]),
		CSS: string(www["/404.css"]),
		Shell: shells.Basic,
		Title: "My App - Not Found",
	}

	pages.Pages["/about"] = &pages.Page{
		Title: "My App - About",
	}

	pages.Pages["/cookies"] = &pages.Page{
		Content: string(www["/cookies.htm"]),
		CSS: string(www["/cookies.css"]),
		Title: "My App - Cookie Policy",
	}

	pages.Pages["/download"] = &pages.Page{
		Content: string(www["/download.htm"]),
		CSS: string(www["/download.css"]),
		Title: "My App - Download",
	}

	pages.Pages["/eula"] = &pages.Page{
		Content: string(www["/eula.htm"]),
		CSS: string(www["/eula.css"]),
		Title: "My App - End User License Agreement (EULA)",
	}

	pages.Pages["/privacy"] = &pages.Page{
		Content: string(www["/privacy.htm"]),
		CSS: string(www["/privacy.css"]),
		Title: "My App - Privacy Policy",
	}

	pages.Pages["/register"] = &pages.Page{
		Content: string(www["/register.htm"]),
		CSS: string(www["/register.css"]),
		Form: "form",
		FormAction: "/download",
		ReCaptchaV2: "6LfgpmoUAAAAAFhnHWF9XHsceqVSFYKH8RDTY-ai",
		Title: "My App - Register",
	}

	pages.Pages["/robots.txt"] = &pages.Page{
		ContentType: "text/plain; charset=utf8",
		Raw: www["/robots.txt"],
	}

	pages.Pages["/terms"] = &pages.Page{
		Content: string(www["/terms.htm"]),
		CSS: string(www["/terms.css"]),
		Title: "My App - Terms of Use",
	}

	var favIcon string
	if raw, ok := www["/favicon.ico"]; ok {
		pages.Pages["/favicon.ico"] = &pages.Page{
			Raw: raw,
		}
		favIcon = "favicon.ico"
	}

	topNav := make(map[string]string)
	topNav["/"] = "Top Page"
	topNav["/about"] = "About"
	topNav["/download"] = "Download"
	topNav["/register"] = "Register"

	for _, k := range []string{
		"",
		"404",
		"/about",
		"/cookies",
		"/download",
		"/eula",
		"/privacy",
		"/register",
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
			if len(p.Mono) < 1 {
				p.Mono = "Source Code Pro"
			}
			if len(p.Sans) < 1 {
				p.Sans = "Noto Sans"
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
