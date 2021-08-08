package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

type GitHook struct {
	Token string `form:"X-Gitlab-Token"`
	Event string `form:"X-Gitlab-Event"`
}

func (c *MainController) Post() {
	// Do some magic here
	if err := c.ParseForm(&u); err != nil {
		//handle error
	}
	c.Abort("403")
}
