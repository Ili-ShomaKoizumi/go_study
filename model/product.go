package model

// 商品
type Product struct {
	ProductId   string // 商品ID
	Title       string // 商品タイトル
	Description string // 説明
	Category    string // 商品カテゴリ
	ImageUrl    string // 商品画像URL
	Price       int    // 値段
	Condition   string // 商品状態(新品 or 中古)
	OwnerId     string // 出品者ID
}

func (product *Product) Create(productId string, title string, desciption string, category string, imageUrl string, price int, condition string, ownerId string) {
	product.ProductId = productId
	product.Title = title
	product.Description = desciption
	product.Category = category
	product.ImageUrl = imageUrl
	product.Price = price
	product.Condition = condition
	product.OwnerId = ownerId
}
