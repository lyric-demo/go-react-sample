package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/LyricTian/lib.v1/mongo"
)

var (
	mHandler *mongo.Handler
)

func init() {
	url := beego.AppConfig.String("MongoURL")
	handler, err := mongo.InitHandler(url)
	if err != nil {
		panic(err)
	}
	mHandler = handler
}
