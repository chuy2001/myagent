package main

import (
	"github.com/astaxie/beego"

	"encoding/json"
	"github.com/gonuts/go-shellquote"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
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

func runCmd(command string) (map[string]string, error) {
	beego.Info("command:" + command)

	_, err := pathExists(command)
	if err != nil {
		beego.Info("file IsNotExist:" + err.Error())
		return nil, err
	}

	split_cmd, err := shellquote.Split(command)

	if err != nil || len(split_cmd) == 0 {
		beego.Info("exec: unable to parse command:" + err.Error())
		return nil, err
	}

	cmd := exec.Command(split_cmd[0], split_cmd[1:]...)
	d, _ := cmd.Output()
	beego.Info("Output:" + string(d))

	if err := json.Unmarshal(d, &xxx); err != nil {
		beego.Info("Unmarshal:" + err.Error())
		return nil, err
	}

	return xxx, nil
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (this *HomeController) Get() {
	beego.Info("InfoTest:" + runtime.GOOS + "  " + this.Ctx.Input.Param(":path"))

	lastname := ""

	if runtime.GOOS == "windows" {
		lastname = ".bat"
	} else {
		lastname = ".sh"
	}

	xxxMap, err := runCmd(this.Ctx.Input.Param(":path") + lastname)

	// xxxMap, err := readFile("roa" + this.Ctx.Input.Param(":path") + "_get.json")
	if err != nil {
		this.Data["json"] = "ROA not Support"
	} else {
		this.Data["json"] = xxxMap
	}
	this.ServeJSON()
}

func main() {
	beego.Router("/*.*", &HomeController{})
	beego.Run()
}
