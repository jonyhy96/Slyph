package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"time"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(123.207.136.18:3306)/k8s?charset=utf8")
	orm.RegisterModel(new(Log))
}

var (
	o    orm.Ormer
	Once sync.Once
)

func Getorg() orm.Ormer {
	Once.Do(func() {
		o = orm.NewOrm()
	})
	return o
}

func Insertlog(log *Log) {
	o := Getorg()
	id, err := o.Insert(log)
	if err != nil {
		fmt.Println(id, err)
	}
}
func Readlog() []Log {
	var logs []Log
	o := Getorg()
	num, err := o.Raw("SELECT time, etype, event FROM log ORDER BY time desc LIMIT 100").QueryRows(&logs)
	if err != nil {
		fmt.Println(num, err)
	}
	return logs
}
func Gettime() string {
	now := time.Now()
	const layout = "Mon Jan _2 15:04:05 MST 2006"
	s := now.Format(layout)
	return s
}
