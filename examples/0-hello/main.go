package main

import (
	"github.com/suite911/cloud911"
	"github.com/suite911/cloud911/easy"
	"github.com/suite911/cloud911/pages"
)

func main() {
	easy.EasyInit()
	pages.Pages["index.html"] = pages.Page{
		Title: "Hello",
		Body: `Hello, world`,
	}
	cloud911.Main()
}
