package global

type IService interface {
	Find(param Parameter) (ret ReturnType)
	FindAll(param Parameter) (ret ReturnType)
	Store(param Parameter) (ret ReturnType)
	Update(param Parameter) (ret ReturnType)
	Delete(param Parameter) (ret ReturnType)
}

type IMakerService interface {
	VerifyNickname(param VerifyNicknameReq) VerifyNicknameRes
	UpdateNickname(param UpdateMakerNicknameReq) UpdateMakerNicknameRes
	Login(param MakerLoginReq) LoginRes
	Register(param RegisterMakerReq) RegisterMakerRes
}

type IMinerService interface {
	Login(param MinerLoginReq) LoginRes
	Register(param RegisterMinerReq) RegisterMinerRes
}
