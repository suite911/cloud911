package shells

import (
	"text/template"

	"github.com/pkg/errors"
)

var Amy *template.Template

func init() {
	text := `<!DOCTYPE html>
<!--[if lte IE 6]><html class="preIE7 preIE8 preIE9"><![endif]-->
<!--[if IE 7]><html class="preIE8 preIE9"><![endif]-->
<!--[if IE 8]><html class="preIE9"><![endif]-->
<!--[if gte IE 9]><!--><html><!--<![endif]-->
<head>
<meta charset="UTF-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
<meta name="viewport" content="width=device-width,initial-scale=1">
{{if .Title}}<title>{{.Title}}</title>
{{end}}{{if .Author}}<meta name="author" content="{{.Author}}">
{{end}}{{if .Description}}<meta name="description" content="{{.Description}}">
{{end}}{{if .Keywords}}<meta name="keywords" content="{{.Keywords}}">
{{end}}{{if .FavIcon}}<link rel="shortcut icon" href="{{.FavIcon}}" type="image/vnd.microsoft.icon">
{{end}}{{if .GoogleFonts}}<link rel="stylesheet" href="//fonts.googleapis.com/css?family={{.GoogleFonts}}" type="text/css">
{{end}}<style type="text/css"><!-- /*<![CDATA[*/
{{if .CSSHead}}{{.CSSHead}}

{{end}}*, *:before, *:after {
	-webkit-box-sizing: inherit;
	-moz-box-sizing: inherit;
	box-sizing: inherit;
}

:root, html, body {
	background-color: #000;
	-webkit-box-sizing: border-box;
	-moz-box-sizing: border-box;
	box-sizing: border-box;
	height: 100%;
	margin: 0;
	padding: 0;
}

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

/* Day Mode */

.page-outer {
	background-color: {{.Vars.Light_Bg}};
	color: {{.Vars.Light_Fg}};
	min-height: 100%;
	padding: 0;
}

.page-inner {
	margin: 0;
	padding: 0 0 {{.Vars.FooterHeight}} 0;
}

header {
	background-color: {{.Vars.Light_Header_Bg}};
	color: {{.Vars.Light_Header_Fg}};
	display: block;
	margin: 0;
	padding: 0;
}

footer {
	background-color: {{.Vars.Light_Footer_Bg}};
	color: {{.Vars.Light_Footer_Fg}};
	display: block;
	font-size: {{.Vars.FooterHeight}};
	height: {{.Vars.FooterHeight}};
	line-height: {{.Vars.FooterHeight}};
	margin: -{{.Vars.FooterHeight}} 0 0 0;
	overflow: hidden;
	text-align: center;
}

.only-day {
	display: inline;
}

.only-night {
	display: none;
}

/* Night Mode */

input[type=checkbox].night:checked + div .only-day {
	display: none;
}

input[type=checkbox].night:checked + div .only-night {
	display: inline;
}

input[type=checkbox].night:checked + div header {
	background-color: {{.Vars.Dark_Header_Bg}};
	color: {{.Vars.Dark_Header_Fg}};
}

input[type=checkbox].night:checked ~ div {
	background-color: {{.Vars.Dark_Bg}};
	color: {{.Vars.Dark_Fg}};
}

input[type=checkbox].night:checked ~ footer {
	background-color: {{.Vars.Dark_Footer_Bg}};
	color: {{.Vars.Dark_Footer_Fg}};
}

/* Inner elements */

div.topnav {
	display: block;
	font-size: {{.Vars.TopNavHeight}};
	height: calc({{.Vars.TopNavHeight}} + 20px);
	line-height: {{.Vars.TopNavHeight}};
	margin: 0;
	overflow: hidden;
	padding: 0;
	text-align: justify;
	width: 100%;
}

.topnavleft {
	border: none;
	display: inline-block;
	height: 100%;
	margin: 0;
	overflow: hidden;
	padding: 0 0;
}

.topnavright {
	border: none;
	display: inline-block;
	margin: 0;
	overflow: hidden;
	padding: 10px 0;
}

.topnavhack {
	border: none;
	display: inline-block;
	margin: 0;
	overflow: hidden;
	padding: 0;
	width: 100%;
}

span.topnav {
	background-color: {{.Vars.Light_TopNav_Bg}};
	border: none;
	border-radius: 0 0 4px 4px;
	color: {{.Vars.Light_TopNav_Fg}};
	display: inline-block;
	height: 100%;
	margin: 0 0px 0 0;
	padding: 10px 4px;
}

span.topnav:hover {
	background-color: {{.Vars.Light_TopNav_Bg_Hover}};
	color: {{.Vars.Light_TopNav_Fg_Hover}};
}

input[type=checkbox].night:checked + div header span.topnav {
	background-color: {{.Vars.Dark_TopNav_Bg}};
	color: {{.Vars.Dark_TopNav_Fg}};
}

input[type=checkbox].night:checked + div header span.topnav:hover {
	background-color: {{.Vars.Dark_TopNav_Bg_Hover}};
	color: {{.Vars.Dark_TopNav_Fg_Hover}};
}

/* The "Night Mode" toggle */
label.night {
	background-color: {{.Vars.Light_Night_Bg}};
	border: 2px solid {{.Vars.Light_Night_Border}};
	border-radius: 4px;
	color: {{.Vars.Light_Night_Fg}};
	display: inline;
	font-size: calc(100% - 2px);
	height: 100%;
	line-height: 100%;
	margin: 0;
	padding: 5px;
}

label.night:hover {
	background-color: {{.Vars.Light_Night_Bg_Hover}};
	border-color: {{.Vars.Light_Night_Border_Hover}};
	color: {{.Vars.Light_Night_Fg_Hover}};
}

input[type=checkbox].night:checked + div header label.night {
	background-color: {{.Vars.Dark_Night_Bg}};
	border-color: {{.Vars.Dark_Night_Border}};
	color: {{.Vars.Dark_Night_Fg}};
}

input[type=checkbox].night:checked + div header label.night:hover {
	background-color: {{.Vars.Dark_Night_Bg_Hover}};
	border-color: {{.Vars.Dark_Night_Border_Hover}};
	color: {{.Vars.Dark_Night_Fg_Hover}};
}

/* Put the associated checkbox over the "Night Mode" text and make it invisible. */
input[type=checkbox].night {
	position: absolute;
	top: 9pt;
	right: 7pt;
	opacity: 0;
}

div.form {
	padding: 16px;
}

hr {
	border: 1px solid #eee;
	margin-bottom: 24px;
}

input[type=text], input[type=password] {
	background-color: {{.Vars.EntryBg}};
	border: 2px solid {{.Vars.EntryBorder}};
	border-radius: 4px;
	color: {{.Vars.EntryFg}};
	display: inline-block;
	margin: 0 0 24px 0;
	outline: none;
	padding: {{.Vars.PaddingWidgetVert}} {{.Vars.PaddingWidgetHorz}};
	width: 100%;
}

input[type=text]:focus, input[type=password]:focus {
	background-color: {{.Vars.FocusBg}};
	border: 2px solid {{.Vars.FocusBorder}};
	color: {{.Vars.FocusFg}};
}

input[type=submit] {
	background-color: {{.Vars.ButtonSubmitBg}};
	border: none;
	border-radius: 4px;
	color: {{.Vars.ButtonSubmitFg}};
	cursor: pointer;
	padding: {{.Vars.PaddingWidgetVert}} {{.Vars.PaddingWidgetHorz}};
	margin: 8px 0;
	font-weight: bold;
	width: 100%;
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
	padding: 0 0 0 2px;
	text-align: left;
}
{{if .CSS}}
{{.CSS}}
{{end}}{{if .CSSTail}}{{.CSSTail}}
{{end}}/*]]>*/ --></style>{{if .Head}}
{{.Head}}{{end}}
{{if .Form}}{{if .ReCaptchaV2}}<script src='https://www.google.com/recaptcha/api.js' async defer></script>
{{end}}<script type="text/javascript"><!-- //<![CDATA[
	function onSubmit(token) {
{{.OnWillSubmit}}
		document.getElementById("{{.Form}}").submit();
{{.OnSubmitted}}
	}
//]]> --></script>
{{end}}</head>
<body>{{.BodyHead}}{{if .Body}}{{.Body}}{{else}}
<input type="checkbox" class="night" id="night" checked />
<div class="page-outer"><div class="page-inner">
	<header class="topnav">
		<div class="topnav">{{.TopNavHead}}
			<div class="topnavleft"{{range $k, $v := .TopNav}}
				><a href="{{$k}}"><span class="topnav">{{$v}}</span></a{{end}}
			></div>
			<div class="topnavright"><label for="night" class="night"><span class="only-day">Lights off &#x263d;</span><span class="only-night">Lights on &#x263c;</span></label></div>
			<div class="topnavhack"></div>{{.TopNavTail}}
		</div>
	</header>
	<header class="header">
		<div class="header">{{.HeaderHead}}{{.Header}}{{.HeaderTail}}
		</div>
	</header>
	<div class="content">{{.ContentHead}}{{if .Form}}
		<div class="form"><form id="{{.Form}}" action="{{if .FormAction}}{{.FormAction}}{{else}}./submit{{end}}" method="POST">{{end}}{{.Content}}{{if .Form}}
		{{if .ReCaptchaV2}}<input type="hidden" id="recaptcha-token" name="recaptcha-token" value="" />
		<input type="submit" id="bsubmit" class="g-recaptcha" data-sitekey="{{.ReCaptchaV2}}"
			value="Submit" data-callback='onSubmit' />
		{{end}}<br /></form></div>{{end}}{{.ContentTail}}
	</div>
</div></div>
<footer class="footer">{{.FooterHead}}{{.Footer}}{{.FooterTail}}
	<div class="copyright">{{.Copyright}}</div>
</footer>
{{end}}{{.BodyTail}}
<script type="text/javascript"><!-- //<![CDATA[
/*{{.DefaultSHA1Implementation}}*/
{{.DefaultCookieStuff}}{{if .JavaScriptHead}}
{{.JavaScriptHead}}
{{end}}{{if .JavaScript}}
{{.JavaScript}}
{{end}}
function onDOMReady(){
{{.OnDOMReady}}
}
function onPageLoaded(){
{{.OnPageLoaded}}
	cookieAgree();
}
if (document.addEventListener) document.addEventListener("DOMContentLoaded", onDOMReady, false);
else if (document.attachEvent) document.attachEvent("onreadystatechange", onDOMReady);
else window.onload = onDOMReady;
if (window.addEventListener) window.addEventListener("load", onPageLoaded, false);
else if (window.attachEvent) window.attachEvent("onload", onPageLoaded);
else window.onload = onPageLoaded;
{{if .JavaScriptTail}}{{.JavaScriptTail}}
{{end}}//]]> --></script>
</body>
</html>
`
	var err error
	if Amy, err = template.New("Amy").Option("missingkey=zero").Parse(text); err != nil {
		panic(errors.Wrap(err, `template.New("Amy").Parse(text)`))
	}
}
