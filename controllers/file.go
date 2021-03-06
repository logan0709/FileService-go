package controllers

import (
	models "FileService/models"

	auth "github.com/abbot/go-http-auth"
	"github.com/astaxie/beego"
)

const DownloadPath = "./upload"

type FileController struct {
	beego.Controller
}

func (this *FileController) Prepare() {
	a := auth.NewBasicAuthenticator("example.com", Secret)
	if username := a.CheckAuth(this.Ctx.Request); username == "" {
		a.RequireAuth(this.Ctx.ResponseWriter, this.Ctx.Request)
	}
}

// 展示页面
func (c *FileController) GetUpload() {
	c.TplName = "upload.html"
}

// 文件上传
func (this *FileController) Upload() {
	f, h, _ := this.GetFile("uploadfile")   //获取上传的文件
	path := DownloadPath + "/" + h.Filename //文件目录
	f.Close()                               //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	this.SaveToFile("uploadfile", path)     //
	this.Redirect("/upload", 302)           //上传成功跳转首页
}

// 文件列表
func (this *FileController) GetList() {
	this.Data["list"] = models.GetFileList(DownloadPath)
	this.TplName = "download.html"
}

// 文件下载
func (this *FileController) Download() {
	filename := this.GetString("filename")
	this.Ctx.Output.Download(DownloadPath + "/" + filename)
}
