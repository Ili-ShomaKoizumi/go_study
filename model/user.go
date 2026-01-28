package model

// ユーザー
type User struct {
	UserId          string // ユーザーID
	UserName        string // ユーザー名
	DisplayName     string // 表示ユーザー名
	Email           string // メールアドレス
	ProfileImageUrl string // プロフィール画像URL
	Description     string // 説明
}

func (user *User) Create(userId string, userName string, displayName string, email string, profileImageUrl string, desciption string) {
	user.UserId = userId
	user.UserName = userName
	user.DisplayName = displayName
	user.Email = email
	user.ProfileImageUrl = profileImageUrl
	user.Description = desciption
}
