package model

type User struct {
	ID          int
	Name        string
	MailAddress string
	Password    string
}

func (u *User) ValidateName() bool {
	return true
}
