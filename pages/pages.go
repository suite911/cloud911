package pages

import (
	"bytes"
	"text/template"

	"github.com/suite911/error911/onfail"
)

// Pages is a map of all of the pages before they are compiled.
var Pages = make(map[string]Page)

// Page is a type representing a page before it is compiled.
type Page struct {
	DefaultCookieStuff string
	
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
	if len(page.DefaultCookieStuff) < 1 {
		page.DefaultCookieStuff = `
function cookieAlert() {
	alert("This site uses cookies to enhance the user experience.");
	return "1"
}
function cookieAgree() {
	if (cookieGet("agreed") == "") {
		cookieSet("agreed", cookieAlert(), 1);
	}
}
function cookieGet(name) {
	var n = name + "=";
	var a = document.cookie.split(';');
	for(var i = 0; i < a.length; i++) {
		var c = a[i];
		while(c.charAt[0] == ' ') {
			c = c.substring(1);
		}
		if (c.indexOf(n) == 0) {
			return c.substring(n.length, c.length);
		}
	}
	return "";
}
function cookieSet(name, value, hours) {
	var nv = name + "=" + value;
	var p = ";path=/";
	if hours === undefined {
		document.cookie = nv + p;
	} else {
		var d = new Date();
		d.setTime(d.getTime() + (hours * 3600000 ));
		var x = ";expires=" + d.toUTCString();
		document.cookie = nv + x + p;
	}
}
		`
	}
	var b bytes.Buffer
	if err := page.Shell.Execute(&b, nil); err != nil {
		return nil, err
	}
	c.Bytes = b.Bytes()
	return c, nil
}
