package model

import "time"

// 購入履歴
type PurHistory struct {
	UserId         string    // ユーザーID
	ProductId      string    // 商品ID
	Title          string    // 商品タイトル
	Price          int       // 金額
	Quantity       int       // 数量
	Category       string    // カテゴリ
	Condition      string    // 商品状態(新品 or 中古)
	DeliveryStatus string    // 配達状況
	OwnerId        string    // 出品者ID
	DtPurchase     time.Time // 購入日時
}

func (purhistory *PurHistory) Create(userId string, productId string, title string, price int, quantity int, category string, condition string, deliveryStatus string, ownerId string) {
	purhistory.UserId = userId
	purhistory.ProductId = productId
	purhistory.Title = title
	purhistory.Price = price
	purhistory.Category = category
	purhistory.Quantity = quantity
	purhistory.Category = category
	purhistory.Condition = condition
	purhistory.DeliveryStatus = deliveryStatus
	purhistory.OwnerId = ownerId
	purhistory.DtPurchase = time.Now()
}
