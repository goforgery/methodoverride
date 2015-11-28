package main

import (
	"github.com/goforgery/forgery2"
	"github.com/goforgery/methodoverride"
)

func main() {
	app := f.CreateApp()
	app.Use(methodoverride.Create())
	app.Listen(3000)
}
