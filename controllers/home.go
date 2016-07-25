package controllers

import ."github.com/Simbory/wemvc"

type Home struct {
	Controller
}

func (ctl Home) GetIndex() ActionResult {
	ctl.ViewData["msg"] = ctl.Session().Get("msg")
	ctl.ViewData["wwwroot"] = RootDir()
	return ctl.View()
}

func (ctl Home) PostIndex() ActionResult {
	msg := ctl.Request.Form.Get("msg")
	ctl.Session().Set("msg", msg)
	ctl.ViewData["msg"] = ctl.Session().Get("msg")
	ctl.ViewData["wwwroot"] = RootDir()
	ctl.ViewData["s"] = ctl.Session().Get("s")
	return ctl.View()
}