package controllers

import (
	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"k8s/models"
	"net/http"
	"runtime/debug"
)

type baseController struct {
	beego.Controller
}

// 200
func (bc *baseController) success(resp interface{}) {
	glog.V(4).Infof("success resp[%+v]\n", resp)
	bc.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	bc.returnJSON(resp)
}

// customError custom code
func (bc *baseController) customError(resp interface{}, statusCode int) {
	glog.Errorln(resp)
	bc.Ctx.ResponseWriter.WriteHeader(statusCode)
	bc.returnJSON(resp)
}

// 500
func (bc *baseController) serviceError(resp interface{}) {
	glog.Errorln(resp)
	bc.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	bc.Data["json"] = resp
	bc.ServeJSON()
	bc.Finish()
}

// 400 参数验证失败
func (bc *baseController) badRequest(resp interface{}) {
	glog.Errorf("Bad Request %+v\n", resp)
	bc.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	bc.returnJSON(resp)
}
func (bc *baseController) catchException() {
	if exp := recover(); exp != nil {
		debug.PrintStack()
		glog.Errorln(string(debug.Stack()))
		glog.Errorln(exp)
		bc.serviceError(models.Error{Error: "panic", Code: 0})
	}
}
func (bc *baseController) returnJSON(resp interface{}) {
	bc.Data["json"] = resp
	bc.ServeJSON()
	bc.Finish()
}
