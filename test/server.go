package main

import (
	"fmt"
	"go/build"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()

	//serve test folder
	m.Use(martini.Static("static"))

	//serve sourcemaps from GOROOT and GOPATH
	m.Use(martini.Static(build.Default.GOROOT, martini.StaticOptions{Prefix: "goroot"}))
	m.Use(martini.Static(build.Default.GOPATH, martini.StaticOptions{Prefix: "gopath"}))
	m.Use(render.Renderer())
	fmt.Println(m)
	m.Run()
}
