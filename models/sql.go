package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
)

func LinkTomysql() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(host)/k8s?charset=utf8")
	if err != nil {
		fmt.Println(err)
		fmt.Println("连接服务器失败")
		glog.Infoln("连接服务器失败！")
	}
	return db, nil
}
func Insertbaseimage(dep string, image string, ver string) error {
	db, err := LinkTomysql()
	defer db.Close()
	if err != nil {
		fmt.Println("insert连接服务器失败")
		fmt.Println(err)
	}
	_, err = db.Exec("insert into dep(dep,image,ver) values(?,?,?)", dep, image, ver)
	return err
}
func Updateimage(dep string, image string) error {
	var s string
	db, err := LinkTomysql()
	rows, err := db.Query("select image from dep where dep=? and ver=?", dep, "new")
	if err != nil {
		return err
	} else {
		for rows.Next() {
			if err := rows.Scan(&s); err != nil {
				return err
			}
		}
	}
	old := "base"
	new := "new"
	if s != image {
		fmt.Println("update")
		_, err := db.Exec("update dep set image=? where dep=? and ver=? ", s, dep, old)
		if err != nil {
			return err
		}
		fmt.Println("update1ed")
		_, err = db.Exec("update dep set image=? where dep=? and ver=? ", image, dep, new)
		if err != nil {
			return err
		}
		fmt.Println("updated")
	}
	return nil
}
func Getbaseimage(dep string) (string, error) {
	var s string
	base := "base"
	db, err := LinkTomysql()
	rows, err := db.Query("select image from dep where dep=? and ver=?", dep, base)
	if err != nil {
		return "", err
	} else {
		for rows.Next() {
			if err := rows.Scan(&s); err != nil {
				return "", err
			}
		}
	}
	fmt.Println(s)
	fmt.Println(dep)
	return s, nil
}
func Truncate() {
	db, err := LinkTomysql()
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.Exec("truncate dep")
	fmt.Println(err)
}
