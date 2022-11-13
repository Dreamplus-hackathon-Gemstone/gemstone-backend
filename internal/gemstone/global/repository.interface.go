package global

type IRepository interface {
	Find(param Parameter) (ret ReturnType)
	FindMany(param Parameter) (ret ReturnType)
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
	FindByContentID(param FindProposalReq) (ret FindProposalRes)
	FindMany(param FindManyProposalReq) (ret FindManyProposalRes)
	Store(param StoreProposalReq) (ret StoreProposalRes)
	Update(param UpdateProposalReq) (ret UpdateProposalRes)
	Delete(param DeleteProposalReq) (ret DeleteProposalRes)
}

type IMovieRepo interface {
	FindByContentID(param FindMovieReq) (ret FindMovieRes)
	FindMany(param FindManyMovieReq) (ret FindManyMovieRes)
	Store(param StoreMovieReq) (ret StoreMovieRes)
}

type ITokenRepo interface {
	Find(param Parameter) (ret ReturnType)
	Store(param Parameter) (ret ReturnType)
}
