package main

import (
	_ "github.com/Simbory/wemvc-sample/controllers"
	_ "github.com/Simbory/wemvc-sample/admin"
	"github.com/Simbory/wemvc"
)

func main() {
	wemvc.
		AddStatic("/css/").
		AddStatic("/js/").
		AddStatic("/favicon.ico").
		Run(8080)
}