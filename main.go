package main

import (
	_"github.com/Simbory/wemvc"
	_"github.com/Simbory/wemvc-sample/controllers"
	_ "github.com/Simbory/wemvc-sample/areas/admin/controllers"
	."github.com/Simbory/wemvc"
)

func main() {
	println("************************************************************")
	println("*   The web application is started...")
	println("************************************************************")
	AppServer.SetStaticPath("/css/")
	AppServer.SetStaticPath("/js/")
	AppServer.SetStaticPath("/favicon.ico")
	AppServer.Run()
	println("************************************************************")
}