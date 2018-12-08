package goods

import (
	"../utils"
)

type GoodsBasic struct {
	Id string
	GoodsName string
	Price int
	Registdate string
}

var goodsInfo = map[int]GoodsBasic {}

func create (goodsName string, price int) {
	id := utils.UniqueId()
	goods_sample := GoodsBasic{id, goodsName, price, ""}
	goodsInfo[len(goodsInfo)] = goods_sample
}

func checkAll (num int, sample chan GoodsBasic) {
	sample <- goodsInfo[num]
}