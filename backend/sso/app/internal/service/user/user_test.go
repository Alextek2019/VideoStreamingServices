package user

import (
	"context"
	"github.com/golang/mock/gomock"
	"testing"
	"vss/sso/internal/domain/user"
	"vss/sso/internal/storage/mocks"
)

func GetUserService(ctrl *gomock.Controller) *Service {
	return &Service{
		repo: mocks.NewMockUserRepo(ctrl),
	}
}

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	srv := GetUserService(ctrl)

	err := srv.Register(context.Background(), user.RegisterUserArgs{})

	if err != nil {
		t.Errorf("cannot register user, error: %v", err.Error())
	}
}

func Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	srv := GetUserService(ctrl)

}

func Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	srv := GetUserService(ctrl)

}

func Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	srv := GetUserService(ctrl)
}
