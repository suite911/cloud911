package vars

import (
	"bytes"
	"text/template"

	"github.com/pkg/errors"
)

var Style1 string

type Style struct {
	ButtonDisabledBg         string
	ButtonDisabledFg         string
	ButtonSubmitBg           string
	ButtonSubmitFg           string
	ButtonSubmitHover        string
	Dark_Bg                  string
	Dark_Fg                  string
	Dark_Footer_Bg           string
	Dark_Footer_Fg           string
	Dark_Header_Bg           string
	Dark_Header_Fg           string
	Dark_Night_Bg_Hover      string
	Dark_Night_Bg            string
	Dark_Night_Border_Hover  string
	Dark_Night_Border        string
	Dark_Night_Fg_Hover      string
	Dark_Night_Fg            string
	Dark_TopNav_Bg_Hover     string
	Dark_TopNav_Bg           string
	Dark_TopNav_Fg_Hover     string
	Dark_TopNav_Fg           string
	EntryBg                  string
	EntryBorder              string
	EntryFg                  string
	FocusBg                  string
	FocusBorder              string
	FocusFg                  string
	FooterHeight             string
	Light_Bg                 string
	Light_Fg                 string
	Light_Footer_Bg          string
	Light_Footer_Fg          string
	Light_Header_Bg          string
	Light_Header_Fg          string
	Light_Night_Bg_Hover     string
	Light_Night_Bg           string
	Light_Night_Border_Hover string
	Light_Night_Border       string
	Light_Night_Fg_Hover     string
	Light_Night_Fg           string
	Light_TopNav_Bg_Hover    string
	Light_TopNav_Bg          string
	Light_TopNav_Fg_Hover    string
	Light_TopNav_Fg          string
	LinkFg                   string
	LinkHover                string
	Mono                     string
	PaddingWidgetHorz        string
	PaddingWidgetVert        string
	Sans                     string
	TopNavHeight             string
}
var StyleTemplate *template.Template

func init() {
	text := `
body, .sans {
	font-family: {{if .Sans}}"{{.Sans}}", {{end}}sans-serif;
}

code, .mono {
	font-family: {{if .Mono}}"{{.Mono}}", {{end}}monospace;
}

a {
	color: {{.LinkFg}};
}

a:hover {
	color: {{.LinkHover}};
}

.page-outer {
	background-color: {{.Light_Bg}};
	color: {{.Light_Fg}};
}

.page-inner {
	padding: 0 0 {{.FooterHeight}} 0;
}

header {
	background-color: {{.Light_Header_Bg}};
	color: {{.Light_Header_Fg}};
}

footer {
	background-color: {{.Light_Footer_Bg}};
	color: {{.Light_Footer_Fg}};
	font-size: {{.FooterHeight}};
	height: {{.FooterHeight}};
	line-height: {{.FooterHeight}};
	margin: -{{.FooterHeight}} 0 0 0;
}

footer a, footer a:active, footer a:focus, footer a:link, footer a:visited {
	color: {{.Light_Footer_Fg}};
}

footer a:hover {
	color: {{.Light_Footer_Fg}};
}

input[type=checkbox].lights-off:checked + div header {
	background-color: {{.Dark_Header_Bg}};
	color: {{.Dark_Header_Fg}};
}

input[type=checkbox].lights-off:checked ~ div {
	background-color: {{.Dark_Bg}};
	color: {{.Dark_Fg}};
}

input[type=checkbox].lights-off:checked ~ footer {
	background-color: {{.Dark_Footer_Bg}};
	color: {{.Dark_Footer_Fg}};
}

input[type=checkbox].lights-off:checked ~ footer a,
input[type=checkbox].lights-off:checked ~ footer a:active,
input[type=checkbox].lights-off:checked ~ footer a:focus,
input[type=checkbox].lights-off:checked ~ footer a:link,
input[type=checkbox].lights-off:checked ~ footer a:visited {
	color: {{.Dark_Footer_Fg}};
}

input[type=checkbox].lights-off:checked ~ footer a:hover {
	color: {{.Dark_Footer_Fg}};
}

/* Inner elements */

div.topnav {
	font-size: {{.TopNavHeight}};
	height: calc({{.TopNavHeight}} + 20px);
	line-height: {{.TopNavHeight}};
}

span.topnav {
	background-color: {{.Light_TopNav_Bg}};
	color: {{.Light_TopNav_Fg}};
}

span.topnav:hover {
	background-color: {{.Light_TopNav_Bg_Hover}};
	color: {{.Light_TopNav_Fg_Hover}};
}

input[type=checkbox].lights-off:checked + div header span.topnav {
	background-color: {{.Dark_TopNav_Bg}};
	color: {{.Dark_TopNav_Fg}};
}

input[type=checkbox].lights-off:checked + div header span.topnav:hover {
	background-color: {{.Dark_TopNav_Bg_Hover}};
	color: {{.Dark_TopNav_Fg_Hover}};
}

/* The "Lights Off" toggle */
label.lights-off {
	background-color: {{.Light_Night_Bg}};
	border: 2px solid {{.Light_Night_Border}};
	color: {{.Light_Night_Fg}};
}

label.lights-off:hover {
	background-color: {{.Light_Night_Bg_Hover}};
	border-color: {{.Light_Night_Border_Hover}};
	color: {{.Light_Night_Fg_Hover}};
}

input[type=checkbox].lights-off:checked + div header label.lights-off {
	background-color: {{.Dark_Night_Bg}};
	border-color: {{.Dark_Night_Border}};
	color: {{.Dark_Night_Fg}};
}

input[type=checkbox].lights-off:checked + div header label.lights-off:hover {
	background-color: {{.Dark_Night_Bg_Hover}};
	border-color: {{.Dark_Night_Border_Hover}};
	color: {{.Dark_Night_Fg_Hover}};
}

input[type=text], input[type=password] {
	background-color: {{.EntryBg}};
	border: 2px solid {{.EntryBorder}};
	color: {{.EntryFg}};
	padding: {{.PaddingWidgetVert}} {{.PaddingWidgetHorz}};
}

input[type=text]:focus, input[type=password]:focus {
	background-color: {{.FocusBg}};
	border: 2px solid {{.FocusBorder}};
	color: {{.FocusFg}};
}

input[type=submit] {
	background-color: {{.ButtonSubmitBg}};
	color: {{.ButtonSubmitFg}};
	padding: {{.PaddingWidgetVert}} {{.PaddingWidgetHorz}};
}

input[type=submit]:hover {
	background-color: {{.ButtonSubmitHover}};
}

input[type=submit]:disabled {
	background-color: {{.ButtonDisabledBg}};
	color: {{.ButtonDisabledFg}};
}

div.copyright {
	font-size: {{.FooterHeight}};
	height: {{.FooterHeight}};
	line-height: {{.FooterHeight}};
	margin: -{{.FooterHeight}} 0 0 0;
}
`
	var err error
	if StyleTemplate, err = template.New("Style").Option("missingkey=zero").Parse(text); err != nil {
		panic(errors.Wrap(err, `template.New("Amy").Parse(text)`))
	}

	var b bytes.Buffer
	if err := StyleTemplate.Execute(&b, Style{
		ButtonDisabledBg:         "",
		ButtonDisabledFg:         "",
		ButtonSubmitBg:           "",
		ButtonSubmitFg:           "",
		ButtonSubmitHover:        "",
		Dark_Bg:                  "#424242",
		Dark_Fg:                  "#fafafa",
		Dark_Footer_Bg:           "#616161",
		Dark_Footer_Fg:           "#fff",
		Dark_Header_Bg:           "#0000",
		Dark_Header_Fg:           "#fff",
		Dark_Night_Bg_Hover:      "#3d5afe",
		Dark_Night_Bg:            "#616161",
		Dark_Night_Border_Hover:  "#3d5afe",
		Dark_Night_Border:        "#616161",
		Dark_Night_Fg_Hover:      "#fff",
		Dark_Night_Fg:            "#f4ff81",
		Dark_TopNav_Bg_Hover:     "#ff1744",
		Dark_TopNav_Bg:           "#b71c1c",
		Dark_TopNav_Fg_Hover:     "#fff",
		Dark_TopNav_Fg:           "#fff",
		EntryBg:                  "#fafafa",
		EntryBorder:              "#000",
		EntryFg:                  "#000",
		FocusBg:                  "#fff",
		FocusBorder:              "#7cf",
		FocusFg:                  "#000",
		FooterHeight:             "14pt",
		Light_Bg:                 "#fff",
		Light_Fg:                 "#000",
		Light_Footer_Bg:          "",
		Light_Footer_Fg:          "",
		Light_Header_Bg:          "",
		Light_Header_Fg:          "",
		Light_Night_Bg_Hover:     "",
		Light_Night_Bg:           "",
		Light_Night_Border_Hover: "",
		Light_Night_Border:       "",
		Light_Night_Fg_Hover:     "",
		Light_Night_Fg:           "",
		Light_TopNav_Bg_Hover:    "",
		Light_TopNav_Bg:          "",
		Light_TopNav_Fg_Hover:    "",
		Light_TopNav_Fg:          "",
		LinkFg:                   "#03a9f4",
		LinkHover:                "#40c4ff",
		Mono:                     "Source Code Pro",
		PaddingWidgetHorz:        "16px",
		PaddingWidgetVert:        "12px",
		Sans:                     "Noto Sans",
		TopNavHeight:             "16pt",
	}); err != nil {
		panic(err)
	}
	Style1 = string(b.Bytes())
}

/* TODO:

		[2]string{"Light_Header_Bg", "#4FC3F7"},//TODO:
		[2]string{"Light_Header_Fg", "#000"},
		[2]string{"Light_TopNav_Bg", "#4FC3F7"},//TODO:
		[2]string{"Light_TopNav_Bg_Hover", "#80D8FF"},//TODO:
		[2]string{"Light_TopNav_Fg", "#fff"},//TODO:
		[2]string{"Light_TopNav_Fg_Hover", "#fff"},//TODO:
		[2]string{"Light_Footer_Bg", "#ccc"},//TODO:
		[2]string{"Light_Footer_Fg", "#000"},
		[2]string{"Light_Night_Border", "#3F51B5"},
		[2]string{"Light_Night_Border_Hover", "#3D5AFE"},
		[2]string{"Light_Night_Bg", "#3F51B5"},
		[2]string{"Light_Night_Bg_Hover", "#3D5AFE"},
		[2]string{"Light_Night_Fg", "#fff"},
		[2]string{"Light_Night_Fg_Hover", "#fff"},


		[2]string{"ButtonCancelBg", "#c00"},//TODO:
		[2]string{"ButtonCancelHover", "#f00"},//TODO:
		[2]string{"ButtonCancelFg", "#fff"},//TODO:
		[2]string{"ButtonSubmitBg", "#03A9F4"},//TODO:
		[2]string{"ButtonSubmitHover", "#40C4FF"},//TODO:
		[2]string{"ButtonSubmitFg", "#fff"},//TODO:
		[2]string{"ButtonDisabledBg", "#777"},//TODO:
		[2]string{"ButtonDisabledFg", "#eee"},//TODO:


*/
