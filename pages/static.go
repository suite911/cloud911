package vars

var FavIcon []ubyte

var Robots = `User-agent: *
Disallow: /api/
`

var Shell = `<!DOCTYPE html>
<!--[if lte IE 6]><html class="preIE7 preIE8 preIE9"><![endif]-->
<!--[if IE 7]><html class="preIE8 preIE9"><![endif]-->
<!--[if IE 8]><html class="preIE9"><![endif]-->
<!--[if gte IE 9]><!--><html><!--<![endif]-->
<head>
	<meta charset="UTF-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
	<meta name="viewport" content="width=device-width,initial-scale=1">
{{if .Title}}	<title>{{.Title}}</title>
{{end}}{{if .Author}}	<meta name="author" content="{{.Author}}">
{{end}}{{if .Description}}	<meta name="description" content="{{.Description}}">
{{end}}{{if .Keywords}}	<meta name="keywords" content="{{.Keywords}}">
{{end}}	<link rel="shortcut icon" href="favicon.ico" type="image/vnd.microsoft.icon">
{{if .GoogleFonts}}	<link rel="stylesheet" href="//fonts.googleapis.com/css?family={{.GoogleFonts}}" type="text/css">
{{end}}	<style type="text/css">
	//
	</style>{{.Head}}
</head>
<body>{{.Body}}
{{if .Autorun}}	<script type="text/javascript">
	function autorun(){{{.Autorun}}}
	if (window.addEventListener) window.addEventListener("load", autorun, false);
	else if (window.attachEvent) window.attachEvent("onload", autorun);
	else window.onload = autorun;
	</script>
{{end}}</body>
</html>
`

var SiteMap string
