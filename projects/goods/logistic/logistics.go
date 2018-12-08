package logistic

import (
	"../utils"
)

type LogisticBasic struct {
	Id string
	GoodsId string
	CityName string
}

var logistics = map[string]string{}

func add (goodsId string, cityName string) string {
	id  := utils.UniqueId()
	/*logistic_sample := logistics{Id:id, goodsId: GoodsId, CityName: cityName}
	logistics[]*/
	return id
}