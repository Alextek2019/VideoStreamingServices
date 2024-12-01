package user

import (
	"github.com/gofrs/uuid"
	domain "vss/sso/internal/domain/user"
	"vss/sso/internal/storage"
)

func UserDTO(user storage.User) domain.User {
	id, _ := uuid.FromString(user.ID)

	return domain.User{
		UserID: id,
		Login:  user.Login,
	}
}
