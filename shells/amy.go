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
*, *:before, *:after {
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

a {
	color: {{.Vars.LinkFg}};
}

a:hover {
	color: {{.Vars.LinkHover}};
}

/* Day Mode */

.page-outer {
	background-color: {{.Vars.BgDay}};
	color: {{.Vars.FgDay}};
	min-height: 100%;
	padding: 0;
}

.page-inner {
	margin: 0;
	padding: 0 0 {{.Vars.FooterHeight}} 0;
}

/* Night Mode */

input[type=checkbox].night:checked ~ div {
	background-color: {{.Vars.BgNight}};
	color: {{.Vars.FgNight}};
}

div.topnav {
	background-color: {{.Vars.TopNavBg1}};
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
	background-color: {{.Vars.TopNavBg}};
	border: none;
	border-radius: 0 0 8px 8px;
	color: {{.Vars.TopNavFg}};
	display: inline-block;
	height: 100%;
	margin: 0 1px 0 0;
	padding: 10px 4px;
}

span.topnav:hover {
	background-color: {{.Vars.TopNavHover}};
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
	background-color: {{.Vars.HeaderBg}};
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

button.submit, button.g-recaptcha {
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

button.submit:hover, button.g-recaptcha:hover {
	background-color: {{.Vars.ButtonSubmitHover}};
}

footer.footer {
	background-color: {{.Vars.FooterBg}};
	margin: -{{.Vars.FooterHeight}} 0 0 0;
	min-height: {{.Vars.FooterHeight}};
	text-align: center;
}
{{.CSS}}/*]]>*/ --></style>{{.Head}}
{{if .Form}}{{if .ReCaptchaV2}}<script src='https://www.google.com/recaptcha/api.js' async defer></script>
{{end}}<script type="text/javascript"><!-- //<![CDATA[
	function onSubmit(token) {
		document.getElementById("{{.Form}}").submit();
	}
//]]> --></script>
{{end}}</head>
<body>{{.BodyHead}}{{if .Body}}{{.Body}}{{else}}
<input type="checkbox" class="night" id="night" checked />
<div class="page-outer"><div class="page-inner">
	<div class="topnav">{{.TopNavHead}}
		<div class="topnavleft"{{range $k, $v := .TopNav}}
			><a href="{{$k}}"><span class="topnav">{{$v}}</span></a{{end}}
		></div>
		<div class="topnavright"><label for="night" class="night">Night Mode &#x263d;</label></div>
		<div class="topnavhack"></div>{{.TopNavTail}}
	</div>
	<div class="header">{{.HeaderHead}}{{.Header}}{{.HeaderTail}}
	</div>
	<div class="content">{{.ContentHead}}{{if .Form}}
		<div class="form"><form id="{{.Form}}" action="{{.FormAction}}" method="POST">{{end}}{{.Content}}{{if .Form}}
		{{if .ReCaptchaV2}}<button class="g-recaptcha" data-sitekey="{{.ReCaptchaV2}}"
			data-callback='onSubmit'>Submit</button>
		{{end}}<br /></form></div>{{end}}{{.ContentTail}}
	</div>
</div></div>
<footer class="footer">{{.FooterHead}}{{.Footer}}{{.FooterTail}}
</footer>
{{end}}{{.BodyTail}}
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
