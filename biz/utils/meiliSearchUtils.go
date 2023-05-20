package utils

import (
	"diploma_search/biz/config"
	"diploma_search/biz/model"
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

var client *meilisearch.Client

func init() {
	client = meilisearch.NewClient(meilisearch.ClientConfig{
		//从yaml读取meili配置
		Host:   "http://" + config.C.Meilisearch.Host + ":" + config.C.Meilisearch.Port,
		APIKey: config.C.Meilisearch.Apikey,
	})
	version, err := client.GetVersion()
	if err != nil {
		log.Fatalln("! ! ! connect error ,place check you the environment.")
	}
	log.Println("connect success , meilisearch version :", version.PkgVersion)
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
	// 改接口有问题，需要改成纯json对象形式
	// log.Println("index : \t", doc.GetIndex(), "title:\t", doc.GetTitle())
	// newdoc, err := json.Marshal(&document{
	// 	Id:     doc.GetDocId(),
	// 	Title:  doc.GetTitle(),
	// 	Genres: doc.GetGenres(),
	// 	Data:   data,
	// })
	// fmt.Println(string(newdoc))
	// if err != nil {
	// 	log.Println("json marshal error !!")
	// }
	//创建文档
	resp, err := client.Index(index).AddDocuments(data)
	if err != nil {
		log.Println("add document error !!")
	}
	log.Println("add document success !!")
	return resp
}
