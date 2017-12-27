package controllers

import (
	"github.com/astaxie/beego"
	"github.com/sethvargo/go-password/password"
	"log"
)

// Password API
type PasswordGeneratorController struct {
	beego.Controller
}

// 默认8位复杂密码
func (this *PasswordGeneratorController) Get() {
	res, err := password.Generate(8, 2, 1, false, false)
	if err != nil {
		log.Fatal(err)
	}
	this.Ctx.WriteString(res)
}

// 自定义密码
func (this *PasswordGeneratorController) Post() {
	len, err := this.GetInt("len")
	if err != nil {
	}
	numDigits, err := this.GetInt("d")
	if err != nil {
	}
	numSymbols, err := this.GetInt("s")
	if err != nil {
	}
	if numDigits < len && numSymbols < len {
		res, err := password.Generate(len, numDigits, numSymbols, false, false)
		if err != nil {
			log.Fatal(err)
		}
		this.Ctx.WriteString(res)
	}
}
