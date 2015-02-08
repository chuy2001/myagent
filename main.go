package main

import (
	"github.com/astaxie/beego"

	"encoding/json"
	"io/ioutil"
)

type HomeController struct {
	beego.Controller
}

var xxx = map[string]string{}

func readFile(filename string) (map[string]string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		beego.Info("ReadFile:" + err.Error())
		return nil, err
	}

	if err := json.Unmarshal(bytes, &xxx); err != nil {
		beego.Info("Unmarshal:" + err.Error())
		return nil, err
	}

	return xxx, nil
}

func (this *HomeController) Get() {
	beego.Info("Info:" + this.Ctx.Input.Param(":path"))

	xxxMap, err := readFile("roa" + this.Ctx.Input.Param(":path") + "_get.json")
	if err != nil {
		beego.Info("ReadFile:" + err.Error())
		this.Data["json"] = "ROA not Support"
		this.ServeJson()
	} else {
		this.Data["json"] = xxxMap
		this.ServeJson()
	}
}

func main() {
	beego.Router("/*.*", &HomeController{})
	beego.Run()
}
