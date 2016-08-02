package controllers

import "github.com/Simbory/wemvc"

type Admin struct {
	wemvc.Controller
}

func (this Admin) Index() wemvc.ActionResult {
	var text string
	ns := this.Namespace()
	if ns != nil {
		text = ns.GetSetting("isDebug")
	} else {
		text = this.Items["name"].(string)
	}
	return this.PlainText("Hello," + text + " " + this.RouteData.ByName("pathInfo"))
}