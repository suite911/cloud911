package shells

import (
	"text/template"

	"github.com/pkg/errors"
)

var Basic *template.Template

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
{{end}}{{if .CSS}}<style type="text/css"><!-- /*<![CDATA[*/
{{.CSS}}
/*]]>*/ --></style>{{end}}{{.Head}}
</head>
<body>{{.BodyHead}}{{if .Body}}{{.Body}}{{else}}
<div class="topnav">{{.TopNavHead}}{{range $k, $v := .TopNav}}
	<a class="topnav" href="{{$k}}"><span class="topnav">{{$v}}</span></a>{{end}}{{.TopNavTail}}
</div>
<div class="header">{{.HeaderHead}}{{.Header}}{{.HeaderTail}}
</div>
<div class="content">{{.ContentHead}}{{.Content}}{{.ContentTail}}
</div>
<div class="footer">{{.FooterHead}}{{.Footer}}{{.FooterTail}}
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
	if Basic, err = template.New("Basic").Parse(text); err != nil {
		panic(errors.Wrap(err, `template.New("Basic").Parse(text)`))
	}
}
