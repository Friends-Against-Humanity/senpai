package domain

type User struct {
	Nickname string `json:"nickname"`
}

func NewUser(nickname string) User {
	return User{
		Nickname: nickname,
	}
}
