package main

import (
	_ "github.com/Simbory/wemvc-sample/controllers"
	_ "github.com/Simbory/wemvc-sample/admin"
	"github.com/Simbory/wemvc"
)

func main() {
	wemvc.
		ServeStaticDir("/css/").
		ServeStaticDir("/js/").
		ServeStaticFile("/favicon.ico").
		Run(8080)
}