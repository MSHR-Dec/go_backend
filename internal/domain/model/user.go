package model

type User struct {
	ID       int
	Name     string
	Password string
	Nickname string
}

func (u User) IsCollectPassword(p string) bool {
	return u.Password == p
}
