package service

import (
	"context"
	"diploma_search/biz/server/rpc"
	"diploma_search/biz/utils"
	"log"

	"github.com/meilisearch/meilisearch-go"
)

type Person struct {
	rpc.UnimplementedPersonServer
}

func (*Person) GetPersons(ctx context.Context, in *rpc.PersonRequest) (*rpc.PesonList, error) {
	//结果为hits里面的interface{}数组
	resp := utils.Search(in.Index, in.KeyWord, meilisearch.SearchRequest{
		Offset: 0,
		Limit:  10,
	})
	// ===================>　转换出现问题
	//类型转换
	perseonList := []*rpc.PersonResponse{}
	for _, item := range resp.Hits {
		//断言
		if personResp, ok := item.(*rpc.PersonResponse); ok {
			perseonList = append(perseonList, personResp)
		} else {
			log.Panicln("type transition error!")
		}
	}

	return &rpc.PesonList{PersonList: perseonList}, nil
}
