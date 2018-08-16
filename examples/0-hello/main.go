package main

import (
	"fmt"

	"github.com/suite911/cloud911"
	"github.com/suite911/cloud911/pages"

	"github.com/suite911/env911"
)

func init() {
	env911.InitAll("MYAPP_", nil, "MyCompany", "myapp")
}

var verbose = false

func main() {
	pages.Pages["index.html"] = pages.Page{
		Title: "Hello",
		Body: `Hello, world`,
	}

	cloud911.Main()
}
