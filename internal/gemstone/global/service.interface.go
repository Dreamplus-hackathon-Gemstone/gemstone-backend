package global

type IService interface {
	Find(param Parameter) (ret ReturnType)
	FindMany(param Parameter) (ret ReturnType)
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

type IProposalService interface {
	Find(param FindProposalReq) (ret FindProposalRes)
	FindMany(param FindManyProposalReq) (ret FindManyProposalRes)
	Store(param StoreProposalReq) (ret StoreProposalRes)
	Update(param UpdateProposalReq) (ret UpdateProposalRes)
	Delete(param DeleteProposalReq) (ret DeleteProposalRes)
}

type IMovieService interface {
	FindByContentID(param FindMovieReq) (ret FindMovieRes)
	FindMany(param FindManyMovieReq) (ret FindManyMovieRes)
	Store(param StoreMovieReq) (ret StoreMovieRes)
}

type IGenreService interface {
	FindMovie(param FindMovieByGenreReq) (ret FindMovieByGenreRes)
	FindProposal(param FindProposalByGenreReq) (ret FindProposalByGenreRes)
	Store(param StoreGenreReq) (ret StoreGenreRes)
}
