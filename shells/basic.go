package shells

var Basic = `<!DOCTYPE html>
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
{{end}}<link rel="shortcut icon" href="favicon.ico" type="image/vnd.microsoft.icon">
{{if .GoogleFonts}}<link rel="stylesheet" href="//fonts.googleapis.com/css?family={{.GoogleFonts}}" type="text/css">
{{end}}{{if .CSS}}<style type="text/css"><!-- /*<![CDATA[*/
{{.CSS}}
/*]]>*/ --></style>{{end}}{{.Head}}
</head>
<body>{{if .Body}}{{.Body}}{{else}}{{.BodyHead}}
<div class="topnav">{{range $k, $v := .TopNav}}
	<a class="topnav" href="{{$k}}"><span class="topnav">{{$v}}</span></a>{{end}}
</div>
<div class="header">{{.Header}}
</div>
<div class="content">{{.Content}}
</div>
<div class="footer">{{.Footer}}
</div>
{{.BodyTail}}{{end}}
<script type="text/javascript"><!-- //<![CDATA[
{{if .JavaScript}}{{.JavaScript}}
{{end}}{{if .OnDOMReady}}function onDOMReady(){{{.OnDOMReady}}
}
{{end}}{{if .OnPageLoaded}}function onPageLoaded(){{{.OnPageLoaded}}
}
{{end}}{{if .OnDOMReady}}if (document.addEventListener) document.addEventListener("DOMContentLoaded", onDOMReady, false);
else if (document.attachEvent) document.attachEvent("onreadystatechange", onDOMReady);
else window.onload = onDOMReady;
{{end}}{{if .OnPageLoaded}}if (window.addEventListener) window.addEventListener("load", onPageLoaded, false);
else if (window.attachEvent) window.attachEvent("onload", onPageLoaded);
else window.onload = onPageLoaded;
{{end}}//]]> --></script>
{{end}}</body>
</html>
`
