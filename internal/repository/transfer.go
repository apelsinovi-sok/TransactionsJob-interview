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
	return user, nil
}

func (r *transferRep) AddMoney(user entity.User, amount int) {
	oldBalance := user.Balance
	newBalance := oldBalance + amount
	user.Balance = newBalance
	r.DB.Save(&user)
}

func (r *transferRep) Transfer(sender entity.User, addressee entity.User, amount int) error {

	if sender.Balance < amount {
		return errors.New("there is not enough money on the balance sheet")
	}

	if sender.Id == addressee.Id {
		return errors.New("the request failed")
	}

	sender.Balance = sender.Balance - amount
	addressee.Balance = addressee.Balance + amount

	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Save(&sender).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Save(&addressee).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error

}
