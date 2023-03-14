package utils

import (
	"diploma_search/biz/config"

	"github.com/meilisearch/meilisearch-go"
)

var client *meilisearch.Client

func init() {
	client = meilisearch.NewClient(meilisearch.ClientConfig{
		//从yaml读取meili配置
		Host:   "http://" + config.C.Meilisearch.Host + ":" + config.C.Meilisearch.Port,
		APIKey: config.C.Meilisearch.Apikey,
	})
}

// 搜索
func Search(index string, keyword string, searchRequest meilisearch.SearchRequest) *meilisearch.SearchResponse {
	searchResponse, err := client.Index(index).Search(keyword, &searchRequest)
	if err != nil {
		panic(err)
	}
	return searchResponse
}

// 增加文档
func addDoc() {

}
