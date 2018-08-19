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
	DefaultCookieStuff, DefaultSHA1Implementation string
	
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

	ProofOfWork int
}

// Compile compiles a page.
func (page *Page) Compile(defaultShell *template.Template, onFail ...onfail.OnFail) (*CompiledPage, error) {
	if len(page.ContentType) < 1 {
		page.ContentType = "text/html"
	}
	c := new(CompiledPage)
	c.ContentType = page.ContentType
	c.ProofOfWork = page.ProofOfWork
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
		page.DefaultSHA1Implementation = `function sha1(msg)
{//(c)2011 x4u https://softwareengineering.stackexchange.com/questions/76939/why-almost-no-webpages-hash-passwords-in-the-client-before-submitting-and-hashi/76947#76947
  function rotl(n,s) { return n<<s|n>>>32-s; };
  function tohex(i) { for(var h="", s=28;;s-=4) { h+=(i>>>s&0xf).toString(16); if(!s) return h; } };
  var H0=0x67452301, H1=0xEFCDAB89, H2=0x98BADCFE, H3=0x10325476, H4=0xC3D2E1F0, M=0x0ffffffff; 
  var i, t, W=new Array(80), ml=msg.length, wa=new Array();
  msg += String.fromCharCode(0x80);
  while(msg.length%4) msg+=String.fromCharCode(0);
  for(i=0;i<msg.length;i+=4) wa.push(msg.charCodeAt(i)<<24|msg.charCodeAt(i+1)<<16|msg.charCodeAt(i+2)<<8|msg.charCodeAt(i+3));
  while(wa.length%16!=14) wa.push(0);
  wa.push(ml>>>29),wa.push((ml<<3)&M);
  for( var bo=0;bo<wa.length;bo+=16 ) {
    for(i=0;i<16;i++) W[i]=wa[bo+i];
    for(i=16;i<=79;i++) W[i]=rotl(W[i-3]^W[i-8]^W[i-14]^W[i-16],1);
    var A=H0, B=H1, C=H2, D=H3, E=H4;
    for(i=0 ;i<=19;i++) t=(rotl(A,5)+(B&C|~B&D)+E+W[i]+0x5A827999)&M, E=D, D=C, C=rotl(B,30), B=A, A=t;
    for(i=20;i<=39;i++) t=(rotl(A,5)+(B^C^D)+E+W[i]+0x6ED9EBA1)&M, E=D, D=C, C=rotl(B,30), B=A, A=t;
    for(i=40;i<=59;i++) t=(rotl(A,5)+(B&C|B&D|C&D)+E+W[i]+0x8F1BBCDC)&M, E=D, D=C, C=rotl(B,30), B=A, A=t;
    for(i=60;i<=79;i++) t=(rotl(A,5)+(B^C^D)+E+W[i]+0xCA62C1D6)&M, E=D, D=C, C=rotl(B,30), B=A, A=t;
    H0=H0+A&M;H1=H1+B&M;H2=H2+C&M;H3=H3+D&M;H4=H4+E&M;
  }
  return tohex(H0)+tohex(H1)+tohex(H2)+tohex(H3)+tohex(H4);
}`
		page.DefaultCookieStuff = `function cookieAlert() {
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
		[2]string{"Dark_Header_Bg", "#444"},//TODO:
		[2]string{"Dark_Header_Fg", "#fff"},//TODO:
		[2]string{"Dark_Footer_Bg", "#444"},//TODO:
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
		[2]string{"ButtonDisabledBg", "#777"},//TODO:
		[2]string{"ButtonDisabledFg", "#eee"},//TODO:
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
