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
<meta charset="UTF-8" />
<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
<meta name="viewport" content="width=device-width,initial-scale=1" />
{{if .PageTitle}}<title>{{.PageTitle}}</title>
{{end}}{{if .Author}}<meta name="author" content="{{.Author}}" />
{{end}}{{if .Description}}<meta name="description" content="{{.Description}}" />
{{end}}{{if .Keywords}}<meta name="keywords" content="{{.Keywords}}" />
{{end}}{{if .FavIcon}}<link rel="shortcut icon" href="{{.FavIcon}}" type="image/vnd.microsoft.icon" />
{{end}}{{if .GoogleFonts}}<link rel="stylesheet" href="//fonts.googleapis.com/css?family={{.GoogleFonts}}" type="text/css" />
{{end}}<link rel="stylesheet" href="//rawgit.com/suite911/cloud911/master/assets/css/amy.css" type="text/css" />
<style type="text/css"><!-- /*<![CDATA[*/
{{if .CSSHead}}{{.CSSHead}}
{{end}}
{{if .CSS}}{{.CSS}}
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
			<div class="topnavright"><label for="night" class="card night"><span class="only-day">Lights off &#x263d;</span><span class="only-night">Lights on &#x263c;</span></label></div>
			<div class="topnavhack"></div>{{.TopNavTail}}
		</div>
	</header>
	<header class="header">
		<div class="header">{{.HeaderHead}}{{.Header}}{{.HeaderTail}}
		</div>
	</header>{{if .ContentTitle}}
	<div class="title">
		<h1>{{.ContentTitle}}</h1>{{if .ContentSubTitle}}
		<p>{{.ContentSubTitle}}</p>{{end}}
		<hr />
	</div>{{end}}
	<div id="content" class="content"{{if .NoScript}} style="display:none"{{end}}>{{.ContentHead}}{{if .Form}}
		<div class="form"><form id="{{.Form}}" action="{{if .FormAction}}{{.FormAction}}{{else}}./submit{{end}}" method="POST">{{end}}{{.Content}}{{if .Form}}
		{{if .ReCaptchaV2}}<input type="hidden" id="recaptcha-token" name="recaptcha-token" value="" />
		<input type="submit" id="bsubmit" class="g-recaptcha" data-sitekey="{{.ReCaptchaV2}}"
			value="Submit" data-callback='onSubmit' />
		{{end}}<br /></form></div>{{end}}{{.ContentTail}}
	</div>{{if .NoScript}}
	<noscript>{{.NoScript}}</noscript>{{end}}
</div></div>
<footer class="footer">{{.FooterHead}}{{.Footer}}{{.FooterTail}}
	<div class="copyright">{{.Copyright}}</div>
</footer>
{{end}}{{.BodyTail}}
<script type="text/javascript"><!-- //<![CDATA[
function hasClass(elem, className) {
	return (' ' + elem.className + ' ').indexOf(' ' + className + ' ') > -1;
}
function replaceState(url) {
	if(typeof history.replaceState === "function") {
		history.replaceState(null, null, url);
		return true
	}
	return false
}
{{.DefaultCookieStuff}}{{if .JavaScriptHead}}
{{.JavaScriptHead}}
{{end}}{{if .JavaScript}}
{{.JavaScript}}
{{end}}
function onDOMReady(){
{{if .NoScript}}
	document.getElementById("content").style.display = "block";
{{end}}{{.OnDOMReady}}
	if(location.hash.length >= 2) {
		var elem = document.getElementById(location.hash.slice(1))
		if(elem) {
			if(hasClass(elem, "fragment-block")) {
				elem.style.display = "block";
			} else if(hasClass(elem, "fragment-inline")) {
				elem.style.display = "inline";
			} else if(hasClass(elem, "fragment-inline-block")) {
				elem.style.display = "inline-block";
			}
			if(!replaceState(location.href.split('#')[0])) {
				location.hash = '';
			}
		}
	} else if(location.href.slice(-1) == '#') {
		replaceState(location.href.slice(0, -1))
	}
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
