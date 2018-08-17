package pages

import (
	"errors"
	"text/template"

	"github.com/suite911/error911/onfail"

	"github.com/valyala/fasthttp"
)

var Pages map[string]Page = make(map[string]Page)

type Page struct {
	Author, Description, FavIcon, Keywords, Title string

	CSS, GoogleWebFonts, Head, JavaScript, OnDOMReady, OnPageLoaded string

	Body, BodyHead, BodyTail, Content, Footer, Header string

	TopNavHead, TopNavTail string

	Shell string

	TopNav map[string]string

	Raw []byte
}

func (page *Page) Serve(ctx *fasthttp.RequestCtx, shell string, onFail ...onfail.OnFail) error {
	fail := func(err error) error {
		ctx.Error("Internal Server Error", 500)
		return onfail.Fail(err, page, onfail.Print, onFail)
	}
	if len(page.Shell) < 1 {
		page.Shell = shell
	}
	if len(page.Raw) > 0 {
		n, err := ctx.Write(page.Raw)
		if err == nil {
			if n == len(page.Raw) {
				return nil
			}
			err = errors.New("Failed to serve complete page")
		}
		return fail(err)
	}
	tmpl, err := template.New("test").Parse(page.Shell)
	if err != nil {
		return fail(err)
	}
	if err = tmpl.Execute(ctx, nil); err != nil {
		return fail(err)
	}
	return nil
}
