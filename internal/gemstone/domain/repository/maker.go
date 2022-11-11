package repository

import (
	"errors"
	"gemstone-backend/internal/gemstone/domain"
	"gemstone-backend/internal/gemstone/global"
	"gorm.io/gorm"
)

type MakerRepo struct {
	db *gorm.DB
}

func NewMakerRepo(db *gorm.DB) global.IMakerRepo {
	return &MakerRepo{db: db}
}

func (m *MakerRepo) TableName() string {
	return "makers"
}
func (m *MakerRepo) FindByNickname(param global.VerifyNicknameReq) (ret global.VerifyNicknameRes) {
	tx := m.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	res := domain.Makers{}
	if err := tx.Table(m.TableName()).Where(domain.Makers{Nickname: param.Nickname}).First(&res); err != nil {
		switch {
		case errors.Is(err.Error, gorm.ErrRecordNotFound):
			return global.VerifyNicknameRes{Success: false, Err: gorm.ErrRecordNotFound}
		default:
			return global.VerifyNicknameRes{Success: false, Err: gorm.ErrInvalidTransaction}
		}
	}
	return global.VerifyNicknameRes{Success: true, Err: nil}
}

func (m *MakerRepo) Store(param global.RegisterMakerReq) (ret global.RegisterMakerRes) {
	//TODO implement me
	panic("implement me")
}

func (m *MakerRepo) Update(param global.UpdateMakerNicknameReq) (ret global.UpdateMakerNicknameRes) {
	//TODO implement me
	panic("implement me")
}
