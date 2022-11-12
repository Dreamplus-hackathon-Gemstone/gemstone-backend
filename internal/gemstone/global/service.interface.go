package global

type IService interface {
	Find(param Parameter) (ret ReturnType)
	FindAll(param Parameter) (ret ReturnType)
	Store(param Parameter) (ret ReturnType)
	Update(param Parameter) (ret ReturnType)
	Delete(param Parameter) (ret ReturnType)
}

type IUserService interface {
	VerifyNickname(param VerifyNicknameReq) VerifyNicknameRes
	UpdateNickname(param UpdateNicknameReq) UpdateNicknameRes
	Login(param LoginReq) LoginRes
	Register(param RegisterReq) RegisterRes
}
