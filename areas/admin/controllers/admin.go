package controllers

import "github.com/Simbory/wemvc"

type Admin struct {
	wemvc.Controller
}

func (this Admin) GetIndex() wemvc.ActionResult {
	return this.PlainText("Hello," + this.Items["name"].(string))
}
