package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	// c.TplName = "index.html"
	this.Redirect("/article", http.StatusFound)
}
