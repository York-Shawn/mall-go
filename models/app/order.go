package app

// 数据库，订单数据映射模型
type Order struct {
	Id              uint64  `gorm:"primaryKey"`
	OpenId          string  `gorm:"open_id"`
	GoodsIdsCount   string  `gorm:"goods_ids_count"`
	GoodsCount      uint    `gorm:"goods_count"`
	TotalPrice      float64 `gorm:"total_price"`
	Status          int     `gorm:"status"` // 订单状态，1-待付款，2-已取消，3-已付款，4-配送中，5-已完成
	Created         string  `gorm:"created"`
	Updated         string  `gorm:"updated"`
	Sid             uint64  `gorm:"sid"`              // 店铺编号
	Name            string  `gorm:"name"`             // 收货人姓名
	Mobile          string  `gorm:"mobile"`           // 手机号
	Province        string  `gorm:"province"`         // 省份
	City            string  `gorm:"city"`             // 城市
	District        string  `gorm:"district"`         // 区（县）
	DetailedAddress string  `gorm:"detailed_address"` // 详细地址
}

// 订单更新参数模型
type OrderUpdateParam struct {
	Id     uint64 `form:"id" binding:"required,gt=0"`
	Status int    `form:"status" binding:"required,gt=0"`
}

// 订单提交参数模型
type OrderSubmitParam struct {
	OpenId  string       `form:"openId" json:"openId" binding:"required"`
	Sid     uint64       `form:"sid" json:"sid" binding:"required,gt=0"`
	Address OrderAddress `form:"address" json:"address"`
}

type OrderAddress struct {
	Id              uint64 `form:"id" json:"id"`
	Name            string `form:"name" json:"name"`
	Mobile          string `form:"mobile" json:"mobile"`
	Province        string `form:"province" json:"province"`
	City            string `form:"city" json:"city"`
	District        string `form:"district" json:"district"`
	DetailedAddress string `form:"detailedAddress" json:"detailedAddress"`
}

// 订单查询参数模型
type OrderQueryParam struct {
	Type   int    `form:"type" json:"type"`
	OpenId string `form:"openId" json:"openId"`
	Sid    uint64 `form:"sid" binding:"required,gt=0"`
}

// 订单列表传输模型
type OrderList struct {
	Id         uint64      `json:"id"`
	Status     int         `json:"status"`
	TotalPrice float64     `json:"totalPrice"`
	GoodsCount uint        `json:"goodsCount"`
	GoodsItem  []GoodsItem `json:"goodsItem"`
}

// 订单商品项传输模型
type GoodsItem struct {
	Id       uint64  `json:"id"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	ImageUrl string  `json:"imageUrl"`
	Count    int     `json:"count"`
}
