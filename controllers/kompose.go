package controllers

import (
	"io"
	"k8s/models"
	"os"
	"os/exec"
)

// AutoBuildController struct
type AutoBuildController struct {
	baseController
}

// AutoDownController struct
type AutoDownController struct {
	baseController
}

// Post to Down
// @Title Post to Down
// @Description Post to login system
// @Success 200 {object} models.Response
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
// @router /down [post]
func (c *AutoDownController) Post() {
	log := models.Log{}
	var backinfo []byte
	var err error
	cmd := exec.Command("kompose", "down")
	if backinfo, err = cmd.Output(); err != nil {
		log.Time = models.Gettime()
		log.Etype = "error"
		log.Event = "Failed to kompose  down " + err.Error()
		go models.Insertlog(&log)
		c.badRequest(models.Error{Error: err.Error()})
		return
	}
	c.success(models.Response{Action: string(backinfo), Status: "success"})
	return
}

// Get to Down
// @Title Get to Down
// @Description Get to login system
// @Success 200 {object} models.Response
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
// @router /down [get]
func (d *AutoDownController) Get() {
	if d.GetSession("username") == "root" {
		d.TplName = "autodown.html"
	} else {
		d.TplName = "404.html"
	}
}

// Get to Build
// @Title Get to Build
// @Description Get to up
// @Success 200 {object} models.Response
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
// @router /up [get]
func (d *AutoBuildController) Get() {
	if d.GetSession("username") == "root" {
		d.TplName = "autoup.html"
	} else {
		d.TplName = "404.html"
	}
}

// Post to Build
// @Title Post to Build
// @Description Post to up
// @Param compose.yaml formData true "docker-compose.yaml"
// @Success 200 {object} models.Response
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
// @router /up [post]
func (d *AutoBuildController) Post() {
	log := models.Log{}
	defer d.catchException()
	var backinfo []byte
	var filename = "./docker-compose.yml"
	var f *os.File
	var err1 error
	requestbody := d.Ctx.Input.CopyBody(1024)
	if len(requestbody) == 0 {
		d.badRequest(models.Error{Error: "文件为空"})
		return
	}
	if checkFileIsExist(filename) { //如果文件存在
		err1 = os.Remove(filename)
		log.Time = models.Gettime()
		log.Etype = "info"
		log.Event = "Already truncate docker-compose.yaml"
		go models.Insertlog(&log)
	}
	f, err1 = os.Create(filename) //创建文件
	Check(err1)
	_, err1 = io.WriteString(f, string(requestbody))
	Check(err1)
	defer f.Close()
	if backinfo, err1 = compose(); err1 != nil {

		d.badRequest(models.Error{Error: err1.Error()})
		return
	}
	d.success(models.Response{Action: string(backinfo), Status: "success"})
	return
}
func Check(e error) {
	log := models.Log{}
	if e != nil {
		log.Time = models.Gettime()
		log.Etype = "error"
		log.Event = e.Error()
		go models.Insertlog(&log)
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
func compose() ([]byte, error) {
	log := models.Log{}
	var backinfo []byte
	var err error
	cmd := exec.Command("kompose", "up")
	if backinfo, err = cmd.Output(); err != nil {
		log.Time = models.Gettime()
		log.Etype = "error"
		log.Event = "Failed to kompose  up " + err.Error()
		go models.Insertlog(&log)
	} else {
		log.Time = models.Gettime()
		log.Etype = "info"
		log.Event = "kompose  up success " + string(backinfo)
		go models.Insertlog(&log)
		return backinfo, err
	}
	return backinfo, err
}
