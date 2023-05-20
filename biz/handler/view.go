package handler

import (
	"context"
	myutils "diploma_search/biz/utils"
	"log"
	"net/http"

	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/meilisearch/meilisearch-go"

	"github.com/cloudwego/hertz/pkg/app"
)

func ViewIndex(c context.Context, ctx *app.RequestContext) {
	ctx.HTML(http.StatusOK, "index.html", utils.H{
		"value":  "Go",
		//随机图
		"imgUrl": "https://random.m2dd.top",
	})
}

// search.html
func ViewSearch(c context.Context, ctx *app.RequestContext) {
	//查询
	keyword := ctx.Query("q")

	resp := myutils.Search("diploma_search", keyword, meilisearch.SearchRequest{
		Offset: 0,
		Limit:  10,
	})
	log.Println("keyword is :\t" + keyword)
	ctx.HTML(http.StatusOK, "search.html", utils.H{
		"keyword": keyword,
		"resp":    resp,
	})
}
