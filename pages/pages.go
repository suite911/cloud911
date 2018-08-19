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
	DefaultCookieStuff string
	
	ContentType string

	Author, Copyright, Description, FavIcon, Keywords, Title string

	CSS, GoogleWebFonts, Head, JavaScript, OnDOMReady, OnPageLoaded string

	Body, BodyHead, BodyTail          string
	Content, ContentHead, ContentTail string
	Footer, FooterHead, FooterTail    string
	Header, HeaderHead, HeaderTail    string

	Form, FormAction, ReCaptchaV2 string

	TopNavHead, TopNavTail string

	TopNav map[string]string

	Vars map[string]string

	Raw []byte

	Shell *template.Template
}

// Compile compiles a page.
func (page *Page) Compile(defaultShell *template.Template, onFail ...onfail.OnFail) (*CompiledPage, error) {
	if len(page.ContentType) < 1 {
		page.ContentType = "text/html"
	}
	c := new(CompiledPage)
	c.ContentType = page.ContentType
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
	if(hours === undefined) {
		document.cookie = nv + p;
	} else {
		var d = new Date();
		d.setTime(d.getTime() + (hours * 3600000 ));
		var x = ";expires=" + d.toUTCString();
		document.cookie = nv + x + p;
	}
}`
	}
	if page.Vars == nil {
		page.Vars = make(map[string]string)
	}
	for _, pair := range [][2]string{
		[2]string{"Dark_Bg", "#2c2f33"},//TODO:
		[2]string{"Dark_Fg", "#fafafa"},//TODO:
		[2]string{"Dark_Header_Bg", "#777"},//TODO:
		[2]string{"Dark_Header_Fg", "#fff"},//TODO:
		[2]string{"Dark_Footer_Bg", "#777"},//TODO:
		[2]string{"Dark_Footer_Fg", "#fff"},//TODO:
		[2]string{"Light_Bg", "#fff"},
		[2]string{"Light_Fg", "#000"},
		[2]string{"Light_Header_Bg", "#ccc"},//TODO:
		[2]string{"Light_Header_Fg", "#000"},
		[2]string{"Light_Footer_Bg", "#ccc"},//TODO:
		[2]string{"Light_Footer_Fg", "#000"},

		[2]string{"PaddingWidgetHorz", "16px"},
		[2]string{"PaddingWidgetVert", "12px"},
		[2]string{"FooterHeight", "16pt"},
		[2]string{"LinkFg", "#03A9F4"},
		[2]string{"LinkHover", "#40C4FF"},
				[2]string{"TopNavBg1", "#0000"},//TODO:
		[2]string{"TopNavBg", "#B0BEC5"},
		[2]string{"TopNavHover", "#E0E0E0"},
		[2]string{"TopNavFg", "#fff"},
		[2]string{"EntryBorder", "#000"},//TODO:
		[2]string{"EntryBg", "#FAFAFA"},//TODO:
		[2]string{"EntryFg", "#000"},//TODO:
		[2]string{"FocusBorder", "#7cf"},//TODO:
		[2]string{"FocusBg", "#fff"},//TODO:
		[2]string{"FocusFg", "#000"},//TODO:
		[2]string{"ButtonCancelBg", "#c00"},//TODO:
		[2]string{"ButtonCancelHover", "#f00"},//TODO:
		[2]string{"ButtonCancelFg", "#fff"},//TODO:
		[2]string{"ButtonSubmitBg", "#03A9F4"},//TODO:
		[2]string{"ButtonSubmitHover", "#40C4FF"},//TODO:
		[2]string{"ButtonSubmitFg", "#fff"},//TODO:
	} {
		if _, ok := page.Vars[pair[0]]; !ok {
			page.Vars[pair[0]] = pair[1]
		}
	}
	var b bytes.Buffer
	if err := page.Shell.Execute(&b, page); err != nil {
		return nil, errors.Wrap(err, "page.Shell.Execute")
	}
	c.Bytes = b.Bytes()
	return c, nil
}
