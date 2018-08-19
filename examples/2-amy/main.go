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
	topNav := make(map[string]string)
	topNav["/"] = "Top Page"
	topNav["/about"] = "About"
	topNav["/register"] = "Register"

	cssHead := string(www["/head.css"])

	var favIcon string
	if raw, ok := www["/favicon.ico"]; ok {
		pages.Pages["/favicon.ico"] = &pages.Page{
			Raw: raw,
		}
		favIcon = "favicon.ico"
	}

	footer := string(www["/footer.htm"])
	googleFonts := "Noto+Sans|Source+Code+Pro"

	pages.Pages[""] = &pages.Page{
		CSSHead: cssHead,
		FavIcon: favIcon,
		Footer: footer,
		GoogleFonts: googleFonts,
		Shell: shells.Amy,
		Title: "My App",
		TopNav: topNav,
	}

	pages.Pages["404"] = &pages.Page{
		Body: string(www["/404.htm"]),
		CSS: string(www["/404.css"]),
		CSSHead: cssHead,
		FavIcon: favIcon,
		Footer: footer,
		GoogleFonts: googleFonts,
		Title: "My App - Not Found",
	}

	pages.Pages["/about"] = &pages.Page{
		CSSHead: cssHead,
		FavIcon: favIcon,
		Footer: footer,
		GoogleFonts: googleFonts,
		Shell: shells.Amy,
		Title: "My App - About",
		TopNav: topNav,
	}

	pages.Pages["/cookies"] = &pages.Page{
		Content: string(www["/cookies.htm"]),
		CSS: string(www["/cookies.css"]),
		CSSHead: cssHead,
		FavIcon: favIcon,
		Footer: footer,
		GoogleFonts: googleFonts,
		Shell: shells.Amy,
		Title: "My App - Cookie Policy",
		TopNav: topNav,
	}

	pages.Pages["/download"] = &pages.Page{
		Content: string(www["/download.htm"]),
		CSS: string(www["/download.css"]),
		CSSHead: cssHead,
		FavIcon: favIcon,
		Footer: footer,
		GoogleFonts: googleFonts,
		Shell: shells.Amy,
		Title: "My App - Download",
		TopNav: topNav,
	}

	pages.Pages["/eula"] = &pages.Page{
		Content: string(www["/eula.htm"]),
		CSS: string(www["/eula.css"]),
		CSSHead: cssHead,
		FavIcon: favIcon,
		Footer: footer,
		GoogleFonts: googleFonts,
		Shell: shells.Amy,
		Title: "My App - End User License Agreement (EULA)",
		TopNav: topNav,
	}

	pages.Pages["/privacy"] = &pages.Page{
		Content: string(www["/privacy.htm"]),
		CSS: string(www["/privacy.css"]),
		CSSHead: cssHead,
		FavIcon: favIcon,
		Footer: footer,
		GoogleFonts: googleFonts,
		Shell: shells.Amy,
		Title: "My App - Privacy Policy",
		TopNav: topNav,
	}

	pages.Pages["/register"] = &pages.Page{
		Content: string(www["/register.htm"]),
		CSSHead: cssHead,
		FavIcon: favIcon,
		Footer: footer,
		Form: "form",
		FormAction: "/download",
		GoogleFonts: googleFonts,
		ReCaptchaV2: "6LfgpmoUAAAAAFhnHWF9XHsceqVSFYKH8RDTY-ai",
		Shell: shells.Amy,
		Title: "My App - Register",
		TopNav: topNav,
	}

	pages.Pages["/robots.txt"] = &pages.Page{
		ContentType: "text/plain; charset=utf8",
		Raw: www["/robots.txt"],
	}

	pages.Pages["/terms"] = &pages.Page{
		Content: string(www["/terms.htm"]),
		CSS: string(www["/terms.css"]),
		CSSHead: cssHead,
		FavIcon: favIcon,
		Footer: footer,
		GoogleFonts: googleFonts,
		Shell: shells.Amy,
		Title: "My App - Terms of Use",
		TopNav: topNav,
	}

	if err := cloud911.Main(); err != nil {
		panic(err)
	}
}

var www = make(map[string][]byte)
