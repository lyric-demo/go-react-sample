package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Article 文章管理
type Article struct {
	ID      bson.ObjectId `json:"ID" bson:"_id,omitempty"`      // ID
	Title   string        `json:"Title" bson:"Title"`           // 标题
	Content string        `json:"Content" bson:"Content"`       // 内容
	Author  string        `json:"Author" bson:"Author"`         // 作者
	IncID   int64         `json:"IncID" bson:"IncID,omitempty"` // 自增ID
}

// Name 集合名称
func (this *Article) Name() string {
	return "Article"
}

// List 查询文章列表
func (this *Article) List() (result []Article, err error) {
	err = mHandler.C(this.Name()).Find(nil).Select(bson.M{"IncID": 0}).Sort("_id").All(&result)
	if err != nil {
		result = make([]Article, 0)
	}
	return
}

// Create 新增文章
func (this *Article) Create() (err error) {
	this.ID = bson.NewObjectId()
	incID, err := mHandler.IncrID(this.Name())
	if err != nil {
		return
	}
	this.IncID = incID
	err = mHandler.C(this.Name()).Insert(this)
	return
}

// UpdateByID 根据ID更新文章信息
func (this *Article) UpdateByID(id bson.ObjectId) error {
	return mHandler.C(this.Name()).UpdateId(id, this)
}

// DeleteByID 根据ID删除文章信息
func (this *Article) DeleteByID(id bson.ObjectId) error {
	return mHandler.C(this.Name()).RemoveId(id)
}

// GetByID 根据ID查询文章信息
func (this *Article) GetByID(id bson.ObjectId) (result Article, err error) {
	err = mHandler.C(this.Name()).FindId(id).One(&result)
	return
}
