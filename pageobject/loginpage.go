package pageobject

import (
	"selenium/base"
)

const (
	username1 string = "login_field"
	pwd1      string = "password"
	submit    string = "commit"
)

type LoginPage struct {
	Base base.Base
}

func (login *LoginPage) InputUserName(username string) {
	err := login.Base.FindEleByID(username1).SendKeys(username)
	if err != nil {
		return
	}
}

func (login *LoginPage) InputPassWord(pwd string) {
	err := login.Base.FindEleByID(pwd1).SendKeys(pwd)
	if err != nil {
		return
	}
}

func (login *LoginPage) ClickLogin() {
	err := login.Base.FindEleByName(submit).Click()
	if err != nil {
		return
	}
}
