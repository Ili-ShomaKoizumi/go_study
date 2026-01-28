package model

// レビュー
type Review struct {
	ProductId string // 商品ID
	Title     string // 商品タイトル
	UserId    string // ユーザーID
	UserName  string // ユーザー名
	SatLevel  int    // 満足レベル
	Comment   string // レビューコメント
}

func (review *Review) Create(productId string, title string, userId string, userName string, satLevel int, comment string) {
	review.ProductId = productId
	review.Title = title
	review.UserId = userId
	review.UserName = userName
	review.SatLevel = satLevel
	review.Comment = comment
}
