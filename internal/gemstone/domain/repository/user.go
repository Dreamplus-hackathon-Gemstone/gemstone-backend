package repository

import (
	"errors"
	"fmt"
	"gemstone-backend/internal/gemstone/domain"
	"gemstone-backend/internal/gemstone/global"
	"gorm.io/gorm"
	"log"
)

type UserRepo struct {
	db *gorm.DB
}

func (u *UserRepo) TableName() string {
	return "users"
}

func NewUserRepo(db *gorm.DB) global.IUserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) FindByNickname(param global.VerifyNicknameReq) (ret global.VerifyNicknameRes) {
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	res := domain.User{}
	if err := tx.Table(u.TableName()).Where(domain.User{Nickname: param.Nickname}).First(&res); err.Error != nil {
		switch {
		case errors.Is(err.Error, gorm.ErrRecordNotFound):
			return global.VerifyNicknameRes{Success: true, Err: fmt.Errorf("possible nickname")}
		default:
			return global.VerifyNicknameRes{Success: false, Err: fmt.Errorf("invalid transaction")}
		}
	}
	return global.VerifyNicknameRes{Success: false, Err: fmt.Errorf("nickname already exists")}
}

func (u *UserRepo) Store(param global.RegisterReq) (ret global.RegisterRes) {
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Create(&domain.User{Nickname: param.Nickname, Name: param.Name, Email: param.Email, Password: param.Password, Account: param.Account}); err.Error != nil {
		log.Printf("Error Rollback : %v", err.Error)
		tx.Rollback()
		return global.RegisterRes{Success: false, Err: err.Error}
	}
	return global.RegisterRes{Success: true, Err: tx.Commit().Error}
}

func (u *UserRepo) Authentication(param global.LoginReq) global.LoginRes {
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	res := domain.User{}
	if err := tx.Table(u.TableName()).Where(&domain.User{Email: param.Email, Password: param.Password}).First(&res); err.Error != nil {
		tx.Rollback()
		return global.LoginRes{Success: false, Err: err.Error}
	}
	return global.LoginRes{Success: true, Err: tx.Commit().Error}
}

func (u *UserRepo) Update(param global.UpdateNicknameReq) (ret global.UpdateNicknameRes) {
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Table(u.TableName()).Where(&domain.User{ID: param.ID}).Update("nickname", param.Nickname); err.Error != nil {
		tx.Rollback()
		return global.UpdateNicknameRes{Success: false, Err: err.Error}
	}
	return global.UpdateNicknameRes{Success: true, Err: tx.Commit().Error}
}
