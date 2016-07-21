package controllers

import "github.com/Simbory/wemvc"

func init() {
	wemvc.AppServer.Route("/", Home{})
	wemvc.AppServer.Route("/download", Download{})
	wemvc.AppServer.Route("/test", TestController{})

	wemvc.AppServer.SetFilter("/", func(ctx wemvc.Context) {
		ctx.SetItem("name", "Simbory")
	})
}