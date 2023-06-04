package service

import (
	"context"
	"diploma_search/biz/server/rpc"
	"diploma_search/biz/utils"
	"encoding/json"
	"log"

	"github.com/meilisearch/meilisearch-go"
)

type Person struct {
	rpc.UnimplementedPersonServer
}

func (*Person) GetPersons(ctx context.Context, in *rpc.PersonRequest) (*rpc.PersonList, error) {
	//结果为hits里面的interface{}数组
	resp := utils.Search(in.Index, in.KeyWord, meilisearch.SearchRequest{
		Offset: 0,
		Limit:  1000,
	})
	log.Println("grpc server call meilisearch ! ")
	data, err := json.Marshal(resp)

	if err != nil {
		log.Panicln("type transition error!")
	}
	return &rpc.PersonList{
		PersonList: data,
	}, nil
}
