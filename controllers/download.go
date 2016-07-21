package controllers

import (
	."github.com/Simbory/wemvc"
	"github.com/Simbory/wemvc/utils"
)

type Download struct {
	Controller
}

func (t Download)Get() ActionResult {
	var file = t.Request.URL.Query().Get("file")
	if len(file) < 1 {
		return t.NotFound()
	}
	file = AppServer.MapPath(file)
	if  !utils.IsFile(file) {
		return t.NotFound()
	}
	return t.File(file, "text/xml")
}