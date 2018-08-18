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
{{end}}{{if .GoogleWebFonts}}<link rel="stylesheet" href="//fonts.googleapis.com/css?family={{.GoogleWebFonts}}" type="text/css">
{{end}}<style type="text/css"><!-- /*<![CDATA[*/
* {
	box-sizing: border-box;
}

:root {
	background-color: #000;

	--padding-widget-horz:	{{.Vars.PaddingWidgetHorz}};
	--padding-widget-vert:	{{.Vars.PaddingWidgetVert}};
	--footer:               {{.Vars.FooterHeight}};
	--a-fg:			{{.Vars.LinkFg}};
	--a-hover:		{{.Vars.LinkHover}};
	--bg-day:		{{.Vars.BgDay}};
	--bg-night:		{{.Vars.BgNight}};
	--bg-topnav:		{{.Vars.TopNavBg1}};
	--bg-header:		{{.Vars.HeaderBg}};
	--bg-footer:		{{.Vars.FooterBg}};
	--fg-day:		{{.Vars.FgDay}};
	--fg-night:		{{.Vars.FgNight}};
	--topnav-bg:		{{.Vars.TopNavBg}};
	--topnav-hover:		{{.Vars.TopNavHover}};
	--topnav-fg:		{{.Vars.TopNavFg}};
	--entry-border:		{{.Vars.EntryBorder}};
	--entry-bg:		{{.Vars.EntryBg}};
	--entry-fg:		{{.Vars.EntryFg}};
	--focus-border:		{{.Vars.FocusBorder}};
	--focus-bg:		{{.Vars.FocusBg}};
	--focus-fg:		{{.Vars.FocusFg}};
	--button-cancel-bg:	{{.Vars.ButtonCancelBg}};
	--button-cancel-hover:	{{.Vars.ButtonCancelHover}};
	--button-cancel-fg:	{{.Vars.ButtonCancelFg}};
	--button-submit-bg:	{{.Vars.ButtonSubmitBg}};
	--button-submit-hover:	{{.Vars.ButtonSubmitHover}};
	--button-submit-fg:	{{.Vars.ButtonSubmitFg}};
}

:root, html, body {
	height: 100%;
	margin: 0;
	padding: 0;
}

a {
	color: var(--a-fg);
}

a:hover {
	color: var(--a-hover);
}

/* Day Mode */

div.wrapper {
	background-color: var(--bg-day);
	color: var(--fg-day);
	min-height: 100%;
	margin: 0 0 calc(0 - var(--footer)) 0;
	padding: 0;
}

/* Night Mode */

input[type=checkbox].night:checked ~ div {
	background-color: var(--bg-night);
	color: var(--fg-night);
}

div.topnav {
	background-color: var(--bg-topnav);
	display: block;
	height: calc(2px + 10px + 15pt + 10px + 2px);
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
	background-color: var(--topnav-bg);
	border: none;
	border-radius: 0 0 8px 8px;
	color: var(--topnav-fg);
	display: inline-block;
	height: 100%;
	margin: 0 1px 0 0;
	padding: 10px 4px;
}

span.topnav:hover {
	background-color: var(--topnav-hover);
}

/* The "Night Mode" toggle */
label.night {
	border: 2px solid #999;
	border-radius: 4px;
	display: inline;
	font-size: 15pt;
	margin: 0;
	padding: 8px;
}

/* Put the associated checkbox over the "Night Mode" text and make it invisible. */
input[type=checkbox].night {
	position: absolute;
	top: 10px;
	right: 10px;
	opacity: 0;
}

.header {
	background-color: var(--bg-header);
}

.container {
	padding: 16px;
}

hr {
	border: 1px solid #eee;
	margin-bottom: 24px;
}

input[type=text], input[type=password] {
	background-color: var(--entry-bg);
	border: 2px solid var(--entry-border);
	border-radius: 4px;
	color: var(--entry-fg);
	display: inline-block;
	margin: 0 0 24px 0;
	outline: none;
	padding: var(--padding-widget-vert) var(--padding-widget-horz);
	width: 100%;
}

input[type=text]:focus, input[type=password]:focus {
	background-color: var(--focus-bg);
	border: 2px solid var(--focus-border);
	color: var(--focus-fg);
}

button.register {
	background-color: var(--button-submit-bg);
	border: none;
	border-radius: 4px;
	color: var(--button-submit-fg);
	cursor: pointer;
	padding: var(--padding-widget-vert) var(--padding-widget-horz);
	margin: 8px 0;
	font-weight: bold;
	width: 100%;
}

button.register:hover {
	background-color: var(--button-submit-hover);
}

footer.footer {
	background-color: var(--bg-footer);
	height: var(--footer);
	text-align: center;
}
{{.CSS}}/*]]>*/ --></style>{{.Head}}
</head>
<body>{{.BodyHead}}{{if .Body}}{{.Body}}{{else}}
<input type="checkbox" class="night" id="night" checked />
<div class="wrapper">
	<div class="topnav">{{.TopNavHead}}
		<div class="topnavleft"{{range $k, $v := .TopNav}}
			><a href="{{$k}}"><span class="topnav">{{$v}}</span></a{{end}}
		></div>
		<div class="topnavright"><label for="night" class="night">Night Mode &#x263d;</label></div>
		<div class="topnavhack"></div>{{.TopNavTail}}
	</div>
	<div class="header">{{.HeaderHead}}{{.Header}}{{.HeaderTail}}
	</div>
	<div class="content">{{.ContentHead}}{{.Content}}{{.ContentTail}}
	</div>
</div>
<div class="footer">
<footer class="footer">{{.FooterHead}}{{.Footer}}{{.FooterTail}}
</footer>
</div>{{end}}{{.BodyTail}}
<script type="text/javascript"><!-- //<![CDATA[
{{.DefaultCookieStuff}}{{if .JavaScript}}{{.JavaScript}}
{{end}}{{if .OnDOMReady}}function onDOMReady(){ {{.OnDOMReady}}
}
{{end}}{{if .OnPageLoaded}}function onPageLoaded(){ {{.OnPageLoaded}}
	cookieAgree();
}
{{end}}{{if .OnDOMReady}}if (document.addEventListener) document.addEventListener("DOMContentLoaded", onDOMReady, false);
else if (document.attachEvent) document.attachEvent("onreadystatechange", onDOMReady);
else window.onload = onDOMReady;
{{end}}{{if .OnPageLoaded}}if (window.addEventListener) window.addEventListener("load", onPageLoaded, false);
else if (window.attachEvent) window.attachEvent("onload", onPageLoaded);
else window.onload = onPageLoaded;
{{end}}//]]> --></script>
</body>
</html>
`
	var err error
	if Amy, err = template.New("Amy").Option("missingkey=zero").Parse(text); err != nil {
		panic(errors.Wrap(err, `template.New("Amy").Parse(text)`))
	}
}
