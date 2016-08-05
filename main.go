package main

import (
	_ "github.com/Simbory/wemvc-sample/controllers"
	_ "github.com/Simbory/wemvc-sample/admin"
	"github.com/Simbory/wemvc"
)

func main() {
	wemvc.
	StaticDir("/css/").
		StaticDir("/js/").
		StaticFile("/favicon.ico").
		Run(8080)
}