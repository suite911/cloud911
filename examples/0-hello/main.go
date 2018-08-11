package main

import (
	"github.com/amy911/snek911/snek"

	"github.com/amy911/srv911/easy"
	"github.com/amy911/srv911/pages"
)

func main() {
	easy.EasyInit()
	pages.Pages["index.html"] = pages.Page{
		Title: "Hello",
		Body: `Hello, world`,
	}
	snek.Main()
}
