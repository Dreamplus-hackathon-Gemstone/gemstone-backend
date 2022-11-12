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

type IUserRepo interface {
	FindByNickname(param VerifyNicknameReq) (ret VerifyNicknameRes)
	Store(param RegisterReq) (ret RegisterRes)
	Update(param UpdateNicknameReq) (ret UpdateNicknameRes)
	Authentication(param LoginReq) (ret LoginRes)
}

type IProposalRepo interface {
	FindByContentID(account Parameter) (ret ReturnType)
	FindMany(param FindManyReq) (ret FindManyProposalRes)
	Store(param Parameter) (ret ReturnType)
	Update(param Parameter) (ret ReturnType)
	Delete(param Parameter) (ret ReturnType)
}

type IMovieRepo interface {
	Find(param FindItemReq) (ret FindMovieRes)
	FindMany(param FindManyReq) (ret ReturnType)
	Store(param Parameter) (ret ReturnType)
	Update(param Parameter) (ret ReturnType)
	Delete(param Parameter) (ret ReturnType)
}

type ITokenRepo interface {
	Find(param Parameter) (ret ReturnType)
	Store(param Parameter) (ret ReturnType)
}
