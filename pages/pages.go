package pages

import (
	"bytes"
	"text/template"

	"github.com/suite911/error911/onfail"

	"github.com/valyala/fasthttp"
)

// Pages is a map of all of the pages before they are compiled.
var Pages = make(map[string]Page)

// Page is a type representing a page before it is compiled.
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

// Compile compiles a page.
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
