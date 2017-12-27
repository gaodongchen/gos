package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
)

// Md5 API
type Md5Controller struct {
	beego.Controller
}

// 生成MD5字符串
func (this *Md5Controller) Post() {
	content := this.Ctx.Input.RequestBody
	ctx := md5.New()
	ctx.Write(content)
	cipher := ctx.Sum(nil)
	res := hex.EncodeToString(cipher)
	this.Ctx.WriteString(res)
}
