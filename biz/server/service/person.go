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
		Limit:  10,
	})
	log.Println("meilisearch : ", resp)
	// ===================>　转换出现问题
	// 类型转换
	// personList := []*rpc.PersonResponse{}
	// for _, item := range resp.Hits {
	// 	fmt.Println(item)
	// 	//断言
	// 	if personResp, ok := item.(model.Person); ok {
	// 		personList = append(personList, &rpc.PersonResponse{
	// 			Id:           personResp.Id,
	// 			Guid:         personResp.Guid,
	// 			UserName:     personResp.UserName,
	// 			Mobile:       personResp.Mobile,
	// 			Address:      personResp.Address,
	// 			Brthday:      personResp.Brthday,
	// 			Email:        personResp.Email,
	// 			Weight:       float32(personResp.Weight),
	// 			Height:       float32(personResp.Height),
	// 			Gender:       personResp.Gender,
	// 			Introduce:    personResp.Introduce,
	// 			Motto:        personResp.Motto,
	// 			Genres:       personResp.Genres,
	// 			Professional: personResp.Professional,
	// 			ProfilePhoto: personResp.ProfilePhoto,
	// 			Title:        personResp.Title,
	// 		})
	// 	} else {
	// 		log.Panicln("type transition error!")
	// 	}
	// }
	data, err := json.Marshal(resp)

	if err != nil {
		log.Panicln("type transition error!")
	}
	// listVal, err := anypb.
	// if err != nil {
	// 	log.Panicln("type transition error!")
	// }
	// personList, err := json.Marshal(resp.Hits)
	// if err != nil {
	// 	log.Panicln("type transition error!")
	// }
	return &rpc.PersonList{
		PersonList: data,
	}, nil
}
