package interfaces

import "test/internal/core/entity"

type TransferRepository interface {
	FindUser(userID int) (entity.User, error)
	AddMoney(user entity.User, amount int)
}
