package pages

import (
	"bytes"
	"strconv"
	"text/template"
	"time"

	"github.com/suite911/error911/onfail"

	"github.com/pkg/errors"
)

// Pages is a map of all of the pages before they are compiled.
var Pages = make(map[string]*Page)

// Page is a type representing a page before it is compiled.
type Page struct {
	ContentType string

	Author, Copyright, Description, FavIcon, Keywords, PageTitle string

	Head, GoogleFonts                          string
	InlineCSS, InlineJavaScript                string
	Body, BodyHead, BodyTail                   string
	Header, HeaderHead, HeaderTail             string
	ContentTitle, ContentSubTitle              string
	Content, ContentHead, ContentTail          string
	Footer, FooterHead, FooterTail             string

	FormAction, ReCaptchaV3 string

	NoScript string

	TopNavHead, TopNavTail string

	TopNav map[string]string

	Raw []byte
	Redirect301 []byte

	Shell *template.Template

	Form bool
}

// Compile compiles a page.
func (page *Page) Compile(defaultShell *template.Template, onFail ...onfail.OnFail) (*CompiledPage, error) {
	if len(page.ContentType) < 1 {
		page.ContentType = "text/html"
	}
	c := new(CompiledPage)
	c.ContentType = page.ContentType
	c.Redirect301 = page.Redirect301
	if len(page.Raw) > 0 {
		c.Bytes = page.Raw
		return c, nil
	}
	if page.Shell == nil {
		page.Shell = defaultShell
	}
	if len(page.Copyright) < 1 {
		page.Copyright = "&copy; " + strconv.Itoa(time.Now().Year())
	}
	var b bytes.Buffer
	if err := page.Shell.Execute(&b, page); err != nil {
		return nil, errors.Wrap(err, "page.Shell.Execute")
	}
	c.Bytes = b.Bytes()
	return c, nil
}
