package main

import (
	_"github.com/Simbory/wemvc"
	_"github.com/Simbory/wemvc-sample/controllers"
	_ "github.com/Simbory/wemvc-sample/areas/admin/controllers"
	"github.com/Simbory/wemvc"
)

func main() {
	wemvc.SetLogFile("log.log")
	wemvc.AddStatic("/css/")
	wemvc.AddStatic("/js/")
	wemvc.AddStatic("/favicon.ico")
	wemvc.Run(8080)
}