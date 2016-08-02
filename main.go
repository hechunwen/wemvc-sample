package main

import (
	_ "github.com/Simbory/wemvc-sample/controllers"
	_ "github.com/Simbory/wemvc-sample/admin"
	"github.com/Simbory/wemvc"
)

func main() {
	wemvc.AddStatic("/css/")
	wemvc.AddStatic("/js/")
	wemvc.AddStatic("/favicon.ico")
	wemvc.Run(8080)
}