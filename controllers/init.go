package controllers

import "github.com/Simbory/wemvc"

func init() {
	wemvc.Route("/", Home{})
	wemvc.Route("/download", Download{})
	wemvc.Route("/test", TestController{})

	wemvc.SetFilter("/", func(ctx wemvc.Context) {
		ctx.SetItem("name", "Simbory")
	})
}