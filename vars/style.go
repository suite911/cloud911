package vars

import (
	"text/template"

	"github.com/pkg/errors"
)

var Style1 string

type Style struct {
	var Mono, Sans string
}
var Template *template.Template

func init() {
	text := `
body, .sans {
	font-family: {{if .Sans}}"{{.Sans}}", {{end}}sans-serif;
}

code, .mono {
	font-family: {{if .Mono}}"{{.Mono}}", {{end}}monospace;
}

a {
	color: {{.Vars.LinkFg}};
}

a:hover {
	color: {{.Vars.LinkHover}};
}

.page-outer {
	background-color: {{.Vars.Light_Bg}};
	color: {{.Vars.Light_Fg}};
}

.page-inner {
	padding: 0 0 {{.Vars.FooterHeight}} 0;
}

header {
	background-color: {{.Vars.Light_Header_Bg}};
	color: {{.Vars.Light_Header_Fg}};
}

footer {
	background-color: {{.Vars.Light_Footer_Bg}};
	color: {{.Vars.Light_Footer_Fg}};
	font-size: {{.Vars.FooterHeight}};
	height: {{.Vars.FooterHeight}};
	line-height: {{.Vars.FooterHeight}};
	margin: -{{.Vars.FooterHeight}} 0 0 0;
}

footer a, footer a:active, footer a:focus, footer a:link, footer a:visited {
	color: {{.Vars.Light_Footer_Fg}};
}

footer a:hover {
	color: {{.Vars.Light_Footer_Fg}};
}

input[type=checkbox].lights-off:checked + div header {
	background-color: {{.Vars.Dark_Header_Bg}};
	color: {{.Vars.Dark_Header_Fg}};
}

input[type=checkbox].lights-off:checked ~ div {
	background-color: {{.Vars.Dark_Bg}};
	color: {{.Vars.Dark_Fg}};
}

input[type=checkbox].lights-off:checked ~ footer {
	background-color: {{.Vars.Dark_Footer_Bg}};
	color: {{.Vars.Dark_Footer_Fg}};
}

input[type=checkbox].lights-off:checked ~ footer a,
input[type=checkbox].lights-off:checked ~ footer a:active,
input[type=checkbox].lights-off:checked ~ footer a:focus,
input[type=checkbox].lights-off:checked ~ footer a:link,
input[type=checkbox].lights-off:checked ~ footer a:visited {
	color: {{.Vars.Dark_Footer_Fg}};
}

input[type=checkbox].lights-off:checked ~ footer a:hover {
	color: {{.Vars.Dark_Footer_Fg}};
}

/* Inner elements */

div.topnav {
	font-size: {{.Vars.TopNavHeight}};
	height: calc({{.Vars.TopNavHeight}} + 20px);
	line-height: {{.Vars.TopNavHeight}};
}

span.topnav {
	background-color: {{.Vars.Light_TopNav_Bg}};
	color: {{.Vars.Light_TopNav_Fg}};
}

span.topnav:hover {
	background-color: {{.Vars.Light_TopNav_Bg_Hover}};
	color: {{.Vars.Light_TopNav_Fg_Hover}};
}

input[type=checkbox].lights-off:checked + div header span.topnav {
	background-color: {{.Vars.Dark_TopNav_Bg}};
	color: {{.Vars.Dark_TopNav_Fg}};
}

input[type=checkbox].lights-off:checked + div header span.topnav:hover {
	background-color: {{.Vars.Dark_TopNav_Bg_Hover}};
	color: {{.Vars.Dark_TopNav_Fg_Hover}};
}

/* The "Lights Off" toggle */
label.lights-off {
	background-color: {{.Vars.Light_Night_Bg}};
	border: 2px solid {{.Vars.Light_Night_Border}};
	color: {{.Vars.Light_Night_Fg}};
}

label.lights-off:hover {
	background-color: {{.Vars.Light_Night_Bg_Hover}};
	border-color: {{.Vars.Light_Night_Border_Hover}};
	color: {{.Vars.Light_Night_Fg_Hover}};
}

input[type=checkbox].lights-off:checked + div header label.lights-off {
	background-color: {{.Vars.Dark_Night_Bg}};
	border-color: {{.Vars.Dark_Night_Border}};
	color: {{.Vars.Dark_Night_Fg}};
}

input[type=checkbox].lights-off:checked + div header label.lights-off:hover {
	background-color: {{.Vars.Dark_Night_Bg_Hover}};
	border-color: {{.Vars.Dark_Night_Border_Hover}};
	color: {{.Vars.Dark_Night_Fg_Hover}};
}

input[type=text], input[type=password] {
	background-color: {{.Vars.EntryBg}};
	border: 2px solid {{.Vars.EntryBorder}};
	color: {{.Vars.EntryFg}};
	padding: {{.Vars.PaddingWidgetVert}} {{.Vars.PaddingWidgetHorz}};
}

input[type=text]:focus, input[type=password]:focus {
	background-color: {{.Vars.FocusBg}};
	border: 2px solid {{.Vars.FocusBorder}};
	color: {{.Vars.FocusFg}};
}

input[type=submit] {
	background-color: {{.Vars.ButtonSubmitBg}};
	color: {{.Vars.ButtonSubmitFg}};
	padding: {{.Vars.PaddingWidgetVert}} {{.Vars.PaddingWidgetHorz}};
}

input[type=submit]:hover {
	background-color: {{.Vars.ButtonSubmitHover}};
}

input[type=submit]:disabled {
	background-color: {{.Vars.ButtonDisabledBg}};
	color: {{.Vars.ButtonDisabledFg}};
}

div.copyright {
	font-size: {{.Vars.FooterHeight}};
	height: {{.Vars.FooterHeight}};
	line-height: {{.Vars.FooterHeight}};
	margin: -{{.Vars.FooterHeight}} 0 0 0;
}
`
	var err error
	if Template, err = template.New("Style").Option("missingkey=zero").Parse(text); err != nil {
		panic(errors.Wrap(err, `template.New("Amy").Parse(text)`))
	}
	Style1, err = Template.execute(Style{
		Mono: "",
		Sans: "",
	})
}
