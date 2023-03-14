// Code generated by hertz generator.

package main

import (
	"context"
	"diploma_search/biz/config"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/network/standard"
)

////go:embed ui/static/* ui/template/*
//var content embed.FS

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
	h := server.Default(
		//读取配置文件端口
		server.WithHostPorts("localhost:"+config.C.Server.Port),
		server.WithMaxRequestBodySize(20<<20),
		server.WithTransport(standard.NewTransporter),
	)
	//自定义解析
	h.Delims("{[{", "}]}")
	h.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})

	h.LoadHTMLGlob("ui/template/views/**")
	//静态文件加载
	h.Static("static", "ui/template")

	h.GET("/raw", func(c context.Context, ctx *app.RequestContext) {
		ctx.HTML(http.StatusOK, "template1.html", map[string]interface{}{
			"now": time.Date(2017, 0o7, 0o1, 0, 0, 0, 0, time.UTC),
		})
	})

	register(h)
	h.Spin()
}
