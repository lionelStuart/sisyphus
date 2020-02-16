package user

type User struct {
	Name string
}

type UserRepository interface {
	Find(id int) (*User, error)
}
