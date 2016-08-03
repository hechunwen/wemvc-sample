package controllers

import "github.com/Simbory/wemvc"

func filterFunc(ctx wemvc.Context) {
	ctx.SetItem("name", "Simbory")
}

func init() {
	wemvc.
		Route("/", Home{}).
		Route("/download", Download{}).
		Route("/test", TestController{}).
		SetFilter("/", filterFunc)
}