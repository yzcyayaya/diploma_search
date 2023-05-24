package handler

import (
	"context"
	"diploma_search/biz/client"
	"diploma_search/biz/server/rpc"
	"log"
	"net/http"
	"strings"

	"github.com/cloudwego/hertz/pkg/common/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

func ViewIndex(c context.Context, ctx *app.RequestContext) {
	ctx.HTML(http.StatusOK, "index.html", utils.H{
		"value": "Go",
		//随机图
		"imgUrl": "https://random.m2dd.top",
	})
}

// search.html
func ViewSearch(c context.Context, ctx *app.RequestContext) {
	//查询
	keyword := ctx.Query("q")
	if keyword == "" || strings.TrimSpace(keyword) == "" {
		log.Println("user input value is null")
		ctx.HTML(http.StatusOK, "index.html", utils.H{
			"value": "请输入值",
			//随机图
			"imgUrl": "https://random.m2dd.top",
		})
	}
	// gprc调用
	resp, err := client.C.GetPersons(c, &rpc.PersonRequest{
		Index:   "diploma_search",
		KeyWord: keyword,
	})
	if err != nil {
		log.Panicln("call grpc is error", err)
	}
	log.Println("keyword is :\t" + keyword)
	ctx.HTML(http.StatusOK, "search.html", utils.H{
		"keyword": keyword,
		"resp":    resp.PersonList,
	})
}
