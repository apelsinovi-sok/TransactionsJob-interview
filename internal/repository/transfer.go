package repository

import (
	"errors"
	"gorm.io/gorm"
	"test/internal/core/entity"
	"test/internal/core/interfaces"
)

type transferRep struct {
	DB *gorm.DB
}

func NewTransferRep(DB *gorm.DB) interfaces.TransferRepository {
	return &transferRep{DB: DB}
}

func (r *transferRep) FindUser(userID int) (entity.User, error) {
	user := entity.User{Id: userID}
	db := r.DB.First(&user)
	if db.RowsAffected == 0 {
		return entity.User{}, errors.New("user is not found")
	}
	//fmt.Println(user.Id)
	//fmt.Println(user.Name)
	return user, nil
}

func (r *transferRep) AddMoney(user entity.User, amount int) {
	oldBalance := user.Balance
	newBalance := oldBalance + amount
	user.Balance = newBalance
	r.DB.Save(&user)
}
