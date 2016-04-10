package controllers

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/astaxie/beego"

	"github.com/LyricTian/go-react-sample/models"
)

// ArticleController 文章管理
type ArticleController struct {
	BaseController
}

// Get
func (this *ArticleController) Get() {
	this.Data["Title"] = "Article List"
	this.TplName = "article/index.html"
	this.Layout = "share/layout.html"
}

// List
func (this *ArticleController) List() {
	var article models.Article
	result, err := article.List()
	if err != nil {
		this.Error(UNKNOWN, "查询文章列表发生异常", err)
		return
	}
	this.Success(result, "查询成功")
}

// GetByID
func (this *ArticleController) GetByID() {
	var article models.Article
	result, err := article.GetByID(bson.ObjectIdHex(this.GetString(":id")))
	if err != nil {
		this.Error(UNKNOWN, "查询文章数据异常!", err)
		return
	}
	this.Success(result, "查询成功")
}

// Create
func (this *ArticleController) Create() {
	this.Data["Title"] = "Create article"
	this.TplName = "article/create.html"
	this.Layout = "share/layout.html"
}

// PostCreate
func (this *ArticleController) PostCreate() {
	var article models.Article
	if err := this.ParseForm(&article); err != nil {
		this.ResParseError(err)
		return
	}
	err := article.Create()
	if err != nil {
		beego.Error(err)
		return
	}
	this.Redirect("/article", http.StatusFound)
}

// Update
func (this *ArticleController) Update() {
	this.Data["ID"] = this.GetString(":id")
	this.Data["Title"] = "Update article"
	this.TplName = "article/update.html"
	this.Layout = "share/layout.html"
}

// PostUpdate
func (this *ArticleController) PostUpdate() {
	var article models.Article
	if err := this.ParseForm(&article); err != nil {
		this.ResParseError(err)
		return
	}
	err := article.UpdateByID(bson.ObjectIdHex(this.GetString(":id")))
	if err != nil {
		beego.Error(err)
		return
	}
	this.Redirect("/article", http.StatusFound)
}

// Delete
func (this *ArticleController) Delete() {
	this.Data["ID"] = this.GetString(":id")
	this.Data["Title"] = "Delete article"
	this.TplName = "article/delete.html"
	this.Layout = "share/layout.html"
}

// PostDelete
func (this *ArticleController) PostDelete() {
	var article models.Article
	err := article.DeleteByID(bson.ObjectIdHex(this.GetString(":id")))
	if err != nil {
		beego.Error(err)
		return
	}
	this.Redirect("/article", http.StatusFound)
}

// View
func (this *ArticleController) View() {
	this.Data["ID"] = this.GetString(":id")
	this.Data["Title"] = "View article"
	this.TplName = "article/view.html"
	this.Layout = "share/layout.html"
}
