package controllers

import "github.com/Simbory/wemvc"

type Admin struct {
	wemvc.Controller
}

func (this Admin) Index() wemvc.ActionResult {
	return this.PlainText("Hello," + this.Items["name"].(string) + " " + this.RouteData.ByName("pathInfo"))
}
