package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"reflect"
	"time"
)

var Session *mgo.Session
var Database *mgo.Database

func init() {
	info := mgo.DialInfo{
		Addrs:    []string{beego.AppConfig.String("mongoUrl")},
		Username: beego.AppConfig.String("mongoUsername"),
		Password: beego.AppConfig.String("mongoPassword"),
		Timeout:  60 * time.Second,
	}
	session, err := mgo.DialWithInfo(&info)
	if err != nil {
		fmt.Println("cannot connect to database: " + err.Error())
		return
	}
	Session = session
	Database = session.DB(beego.AppConfig.String("database"))
}

func Contains(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}

	return false
}