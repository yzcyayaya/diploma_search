package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func Test_Readjson(T *testing.T) {
	file, _ := os.OpenFile("/home/m2/Downloads/shga_sample_750k/person_info.json", os.O_CREATE|os.O_RDWR, 0666)
	defer file.Close()
	//创建map，用于接收解码好的数据
	dataMap1 := make(map[string]interface{})
	//创建文件的解码器
	decoder := json.NewDecoder(file)
	//解码文件中的数据，丢入dataMap所在的内存
	err8 := decoder.Decode(&dataMap1)
	if err8 != nil {
		fmt.Println(err8)
		return
	}
	// fmt.Println("解码成功", dataMap1)
	fmt.Println(dataMap1["_id"])
}
