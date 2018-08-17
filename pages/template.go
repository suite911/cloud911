package pages

import (
	"bytes"
	"text/template"

	"github.com/suite911/error911/onfail"

	"github.com/valyala/fasthttp"
)

type CompiledPage struct {
	ContentType string
	Bytes       []byte
}

func (c *CompiledPage) Serve(ctx *fasthttp.RequestCtx) {
	if len(c.ContentType) > 0 {
		ctx.SetContentType(c.ContentType)
	}
	ctx.Write(c.Bytes)
}

var CompiledPages = make(map[string]*CompiledPage)

func Compile(defaultShell *template.Template, onFail ...onfail.OnFail) error {
	for k, v := range Pages {
		c, err := v.Compile(defaultShell, onFail...)
		if err != nil {
			return err
		}
		CompiledPages[k] = c
	}
	return nil
}

var Pages = make(map[string]Page)

type Page struct {
	ContentType string

	Author, Description, FavIcon, Keywords, Title string

	CSS, GoogleWebFonts, Head, JavaScript, OnDOMReady, OnPageLoaded string

	Body, BodyHead, BodyTail          string
	Content, ContentHead, ContentTail string
	Footer, FooterHead, FooterTail    string
	Header, HeaderHead, HeaderTail    string

	TopNavHead, TopNavTail string

	TopNav map[string]string

	Raw []byte

	Shell *template.Template
}

func (page *Page) Compile(defaultShell *template.Template, onFail ...onfail.OnFail) (*CompiledPage, error) {
	c := new(CompiledPage)
	c.ContentType = page.ContentType
	if len(page.Raw) > 0 {
		c.Bytes = page.Raw
		return c, nil
	}
	if len(page.ContentType) < 1 {
		page.ContentType = "text/html"
	}
	if page.Shell == nil {
		page.Shell = defaultShell
	}
	var b bytes.Buffer
	if err := page.Shell.Execute(&b, nil); err != nil {
		return nil, err
	}
	c.Bytes = b.Bytes()
	return c, nil
}
