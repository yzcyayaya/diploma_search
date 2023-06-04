package utils

import (
	"diploma_search/biz/config"
	"diploma_search/biz/model"
	"flag"
	"log"

	"github.com/meilisearch/meilisearch-go"
)

//meilsearch通过改字段去新增
// type document struct {
// 	Id     string      `json:"id"`
// 	Title  string      `json:"title"`
// 	Genres []string    `json:"genres"`
// 	Data   interface{} `json:"data"`
// }

var (
	client  *meilisearch.Client
	hostUrl string
)

func init() {
	//没有flage则从yaml读取meili配置
	flag.StringVar(&hostUrl, "meili_url", config.C.Meilisearch.Host+":"+config.C.Meilisearch.Port, "meilisearch: meilisearch:7700")
	flag.Parse()
	client = meilisearch.NewClient(meilisearch.ClientConfig{

		Host:   "http://" + hostUrl,
		APIKey: config.C.Meilisearch.Apikey,
	})
	version, err := client.GetVersion()
	if err != nil {
		log.Fatalln("! ! ! connect error ,place check you the environment.")
	}
	log.Println("connect success :"+hostUrl+", meilisearch version :", version.PkgVersion)
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
func AddDoc[T model.Person](index string, data []T) *meilisearch.TaskInfo {
	//创建文档
	_, err := client.CreateIndex(&meilisearch.IndexConfig{
		Uid:        index,
		PrimaryKey: "id",
	})
	if err != nil {
		log.Println("index already existing !!")
	}
	resp, err := client.Index(index).AddDocuments(data)
	if err != nil {
		log.Println("add document error !!")
	}
	log.Println("add document success !!")
	return resp
}
