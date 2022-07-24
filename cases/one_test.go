package cases

import (
	"selenium/pageobject"
	"testing"
)

func TestOne(t *testing.T) {
	login := pageobject.LoginPage{Base: page}
	login.InputUserName("1418500295@qq.com")
	login.InputPassWord("gxm94412")
	login.ClickLogin()
}
