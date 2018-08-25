package main

import (
	"github.com/suite911/cloud911"
	"github.com/suite911/cloud911/pages"

	"github.com/suite911/env911"
)

func init() {
	env911.InitAll("MYAPP_", nil, "MyCompany", "myapp")
}

var verbose = false

func main() {
	pages.Pages["/"] = &pages.Page{
		PageTitle: "Hello",
		Body: `Hello, world`,
	}

	if err := cloud911.Main(); err != nil {
		panic(err)
	}
}
