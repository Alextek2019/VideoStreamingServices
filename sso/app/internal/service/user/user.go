package user

import "context"

type UserService struct {
}

func New(ctx context.Context) *UserService {
	return &UserService{}
}

func (u *UserService) RegisterUser() {

}
