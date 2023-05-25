package handler

import (
	"context"
	"diploma_search/biz/client"
	"diploma_search/biz/server/rpc"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/meilisearch/meilisearch-go"

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
	// gprc调用 ===>>>解码有问题
	resp, err := client.C.GetPersons(c, &rpc.PersonRequest{
		Index:   "diploma_search",
		KeyWord: keyword,
	})
	log.Println("resp : ", resp.PersonList)
	if err != nil {
		log.Panicln("call grpc is error", err)
	}
	personList := struct {
		PersonList *meilisearch.SearchResponse
	}{
		PersonList: &meilisearch.SearchResponse{},
	}
	json.Unmarshal(resp.PersonList, &personList)
	log.Println("keyword is :\t" + keyword)

	log.Println("resp : ", &personList)
	ctx.HTML(http.StatusOK, "search.html", utils.H{
		"keyword": keyword,
		"resp":    personList.PersonList,
	})
}
