package global

type IRepository interface {
	Find(param Parameter) (ret ReturnType)
	FindAll(param Parameter) (ret ReturnType)
	Store(param Parameter) (ret ReturnType)
	Update(param Parameter) (ret ReturnType)
	Delete(param Parameter) (ret ReturnType)
}

type IGenreRepo interface {
	FindByGenre(param Parameter) (ret ReturnType)
	Store(param Parameter) (ret ReturnType)
}

type IMakerRepo interface {
	FindByNickname(nickname VerifyNicknameReq) (ret VerifyNicknameRes)
	Store(param RegisterMakerReq) (ret RegisterMakerRes)
	Update(param UpdateMakerNicknameReq) (ret UpdateMakerNicknameRes)
}

type IMinerRepo interface {
	Find(account FindMinerReq) (ret FindMinerRes)
	Store(param RegisterMinerReq) (ret RegisterMinerRes)
}

type IProposalRepo interface {
	FindByContentID(account FindMinerReq) (ret FindMinerRes)
	FindMany(param FindManyReq) (ret FindManyProposalRes)
	Store(param Parameter) (ret ReturnType)
	Update(param Parameter) (ret ReturnType)
	Delete(param Parameter) (ret ReturnType)
}

type IMovieRepo interface {
	Find(account FindItemReq) (ret FindMovieRes)
	FindMany(param FindManyReq) (ret ReturnType)
	Store(param Parameter) (ret ReturnType)
	Update(param Parameter) (ret ReturnType)
	Delete(param Parameter) (ret ReturnType)
}

type ITokenRepo interface {
	Find(account Parameter) (ret ReturnType)
	Store(param Parameter) (ret ReturnType)
}
