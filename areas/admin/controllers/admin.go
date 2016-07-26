package controllers

import "github.com/Simbory/wemvc"

type Admin struct {
	wemvc.Controller
}

func (ctrl Admin) GetIndex() wemvc.ActionResult {
	return ctrl.PlainText("Hello," + ctrl.Items["name"].(string) + ctrl.Request.URL.String())
}
