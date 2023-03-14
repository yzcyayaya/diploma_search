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
		"imgUrl": "https://api.mtyqx.cn/api/random.php",
	})
}

// search.html
func ViewSearch(c context.Context, ctx *app.RequestContext) {
	//查询
	keyword := ctx.Query("q")	

	resp := myutils.Search("test", keyword, meilisearch.SearchRequest{})

	log.Println("keyword is :\t" + keyword)
	ctx.HTML(http.StatusOK, "search.html", utils.H{
		"title": keyword,
		"resp":  resp,
	})
}
