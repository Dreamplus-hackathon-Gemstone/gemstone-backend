package service

import "gemstone-backend/internal/gemstone/global"

type UserService struct {
	UserRepo global.IUserRepo
}

func NewUserService(userRepo global.IUserRepo) global.IUserService {
	return &UserService{UserRepo: userRepo}
}

func (u *UserService) VerifyNickname(param global.VerifyNicknameReq) global.VerifyNicknameRes {
	return u.UserRepo.FindByNickname(param)
}

func (u *UserService) UpdateNickname(param global.UpdateNicknameReq) global.UpdateNicknameRes {
	return u.UserRepo.Update(param)
}

func (u *UserService) Login(param global.LoginReq) global.LoginRes {
	return u.UserRepo.Authentication(param)
}

func (u *UserService) Register(param global.RegisterReq) global.RegisterRes {
	return u.UserRepo.Store(param)
}
