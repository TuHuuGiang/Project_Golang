package repo

type UserRepo struct{}

func NewUserService() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetInfoUser() string {
	return "Golang"
}
