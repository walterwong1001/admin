package services

import repository "github.com/weitien/admin/repositories"

type UserService struct {
	Repository *repository.userRepositoryImpl
}

func NewUserService() *UserService {
	return &UserService{
		Repository: repository.NewUserRepository(),
	}
}

func (s *UserService) CreateUser() {

}
